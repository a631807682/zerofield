package scopes

import (
	"reflect"

	"gorm.io/gorm"
)

type config struct {
	// allow update zero field. include all field if empty.
	Includes []string
}

// UpdateZeroFields allow update zero cloumns which specified.
// just work for db.Updates(&model) and db.Save(&model).
// if cloumns is empty, all field will be save like db.Select("*"")
func UpdateZeroFields(cloumns ...string) func(db *gorm.DB) *gorm.DB {
	return update(&config{
		Includes: cloumns,
	})
}

func update(cfg *config) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		stmt := db.Statement
		if len(stmt.Selects) > 0 {
			// do not change when a field is specified
			return db
		}

		if stmt.Dest != nil {
			reflectValue := reflect.Indirect(reflect.ValueOf(stmt.Dest))
			for reflectValue.Kind() == reflect.Ptr {
				reflectValue = reflectValue.Elem()
			}
			if reflectValue.Kind() != reflect.Struct {
				// not support other dest type
				return db
			}

			if len(cfg.Includes) == 0 {
				db.Select("*")
			} else {
				db.Set("zerofield:includes", cfg.Includes)
				return db
			}
		}
		return db
	}
}

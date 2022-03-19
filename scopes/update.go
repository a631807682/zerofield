package scopes

import (
	"reflect"

	"gorm.io/gorm"
)

type config struct {
	// allow update zero field. include all field if empty.
	Includes []string
}

// UpdateFields allow update zero cloumns which specified.
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
				replaceSelects(db, cfg.Includes)
				return db
			}
		}
		return db
	}
}

func replaceSelects(db *gorm.DB, includes []string) {
	updateProcessor := db.Callback().Update()
	beforeUpdateHandler := updateProcessor.Get("gorm:before_update")

	// replace selects after before update hooks
	var replaceBeforeUpdateHandler = func(handler func(*gorm.DB), includes []string) func(*gorm.DB) {
		return func(tx *gorm.DB) {
			defer updateProcessor.Replace("gorm:before_update", handler)
			handler(tx)

			includeFieldMap := make(map[string]bool)
			for _, fname := range includes {
				includeFieldMap[fname] = true
			}

			stmt := tx.Statement
			reflectValue := stmt.ReflectValue
			selectColumns := make([]string, 0)
			for _, f := range stmt.Schema.Fields {
				_, isZero := f.ValueOf(stmt.Context, reflectValue)
				if includeFieldMap[f.Name] || !isZero {
					selectColumns = append(selectColumns, f.Name)
				}
			}
			stmt.Selects = selectColumns
		}
	}

	updateProcessor.Replace("gorm:before_update",
		replaceBeforeUpdateHandler(beforeUpdateHandler, includes),
	)
}

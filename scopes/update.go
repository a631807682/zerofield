package scopes

import (
	"gorm.io/gorm"
)

type Config struct {
	Associations bool
}

var DefaultCfg = &Config{
	Associations: false,
}

// UpdateScopes allow update zero field.
// Only work for db.Updates(&model).
func UpdateScopes(c ...*Config) func(db *gorm.DB) *gorm.DB {
	cfg := DefaultCfg
	if len(c) > 0 {
		cfg = c[0]
	}

	return func(db *gorm.DB) *gorm.DB {
		stmt := db.Statement
		if cfg.Associations && len(stmt.Selects) == 1 && stmt.Selects[0] == "*" {
			// if s := stmt.Schema; s != nil && len(s.Fields) > 0 {
			// 	for _, field := range s.Fields {
			// 		// selected := selectedColumns[field.DBName] || selectedColumns[field.Name]
			// 		// if selected || (!restricted && field.Readable) {
			// 		// 	if v, isZero := field.ValueOf(stmt.Context, reflectValue); !isZero || selected {
			// 		// 		if field.DBName != "" {
			// 		// 			conds = append(conds, clause.Eq{Column: clause.Column{Table: clause.CurrentTable, Name: field.DBName}, Value: v})
			// 		// 		} else if field.DataType != "" {
			// 		// 			conds = append(conds, clause.Eq{Column: clause.Column{Table: clause.CurrentTable, Name: field.Name}, Value: v})
			// 		// 		}
			// 		// 	}
			// 		// }
			// 	}
			// }
			// relation empty

			return db
		} else if len(stmt.Selects) > 0 {
			// do not change when selects not empty
			return db
		}

		if stmt.Dest != nil {
			// do not change when a field is specified
			switch stmt.Dest.(type) {
			case map[string]interface{},
				*map[string]interface{},
				[]map[string]interface{},
				*[]map[string]interface{}:
				return db
			}

			db.Select("*")
		}
		return db
	}
}

package plugins

import (
	"gorm.io/gorm"
)

// ZeroFieldPlugin gorm plugin
type ZeroFieldPlugin struct{}

// Name gorm plugin name
func (*ZeroFieldPlugin) Name() string {
	return "gorm:zerofield"
}

// Initialize gorm plugin initialize
func (*ZeroFieldPlugin) Initialize(db *gorm.DB) error {
	updateProcessor := db.Callback().Update()
	register := updateProcessor.Before("gorm:update").After("gorm:save_before_associations")

	register.Register("zerofield:force_update_field", func(tx *gorm.DB) {
		stmt := tx.Statement
		if len(stmt.Selects) > 0 {
			return
		}

		if val, ok := tx.Get("zerofield:includes"); ok {
			if includes, ok := val.([]string); ok && len(includes) > 0 {
				includeFieldMap := make(map[string]bool)
				for _, fname := range includes {
					includeFieldMap[fname] = true
				}

				stmt := tx.Statement
				reflectValue := stmt.ReflectValue
				selectColumns := make([]string, 0)
				for _, f := range stmt.Schema.Fields {
					if f.Updatable {
						_, isZero := f.ValueOf(stmt.Context, reflectValue)
						if includeFieldMap[f.Name] || !isZero {
							selectColumns = append(selectColumns, f.Name)
						}
					}
				}
				stmt.Selects = selectColumns
			}
		}
	})
	return nil
}

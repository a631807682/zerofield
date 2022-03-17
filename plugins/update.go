package plugins

import (
	"gorm.io/gorm"
	gut "gorm.io/gorm/utils"
)

// `gorm:zf_force_update:true`
const ForceUpdateTag = "ZF_FORCE_UPDATE"

type ZeroFieldPlugin struct{}

func (s *ZeroFieldPlugin) Name() string {
	return "gorm:zerofield"
}

func (s *ZeroFieldPlugin) Initialize(db *gorm.DB) error {
	updateProcessor := db.Callback().Update()
	register := updateProcessor.Before("gorm:update").After("gorm:save_before_associations")

	register.Register("zerofield:force_update_field", func(tx *gorm.DB) {
		stmt := tx.Statement
		if len(stmt.Selects) > 0 {
			return
		}

		selectColumns := make([]string, 0)
		for _, f := range stmt.Schema.Fields {
			_, isZero := f.ValueOf(stmt.Context, stmt.ReflectValue)
			if !isZero || gut.CheckTruth(f.TagSettings[ForceUpdateTag]) {
				selectColumns = append(selectColumns, f.Name)
			}
		}
		stmt.Selects = selectColumns
	})
	return nil
}

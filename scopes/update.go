package zerofield

import (
	"gorm.io/gorm"
)

type Config struct {
}

func UpdateScope(c *Config) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// org := r.Query("org")

		// if org != "" {
		// 	var organization Organization
		// 	if db.Session(&Session{}).First(&organization, "name = ?", org).Error == nil {
		// 		return db.Where("org_id = ?", organization.ID)
		// 	}
		// }

		// db.AddError("invalid organization")
		return db
	}
}

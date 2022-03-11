package zerofield

import (
	"gorm.io/gorm"
)

type Config struct {
}

func UpdateScope(c *Config) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

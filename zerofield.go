package zerofield

import (
	"github.com/a631807682/zerofield/plugins"
	"github.com/a631807682/zerofield/scopes"
	"gorm.io/gorm"
)

// UpdateZeroFields allow update zero cloumns which specified.
// just work for db.Updates(&model) and db.Save(&model).
// if cloumns is empty, all field will be save like db.Select("*"")
func UpdateScopes(cloumns ...string) func(db *gorm.DB) *gorm.DB {
	return scopes.UpdateZeroFields(cloumns...)
}

// NewPlugin gorm plugin for zero value field
func NewPlugin() *plugins.ZeroFieldPlugin {
	return &plugins.ZeroFieldPlugin{}
}

package zerofield

import (
	"github.com/a631807682/zerofield/plugins"
	"github.com/a631807682/zerofield/scopes"
)

// UpdateScopes update zero value field by cloumns
// when cloumns is empty, update all
var UpdateScopes = scopes.UpdateZeroFields

// NewPlugin gorm plugin for zero value field
// force update field when defined tage in mode
// like:
// type Foo struct{
// 	Bar string `gorm:zf_force_update:true`
// }
func NewPlugin() *plugins.ZeroFieldPlugin {
	return &plugins.ZeroFieldPlugin{}
}

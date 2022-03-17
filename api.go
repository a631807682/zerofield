package zerofield

import (
	"zerofield/plugins"
	"zerofield/scopes"
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

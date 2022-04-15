# zerofield

gorm plugin for allow update zero value field.

[![go report card](https://goreportcard.com/badge/github.com/a631807682/zerofield "go report card")](https://goreportcard.com/report/github.com/a631807682/zerofield)
[![test status](https://github.com/a631807682/zerofield/workflows/tests/badge.svg?branch=master "test status")](https://github.com/a631807682/zerofield/actions)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/a631807682/zerofield)
![visitor badge](https://visitor-badge.glitch.me/badge?page_id=a631807682.zerofield)

# Desc

> When update with struct, GORM will only update non-zero fields, you might want to use map to update attributes or use Select to specify fields to update
> [Updates-multiple-columns](https://gorm.io/docs/update.html#Updates-multiple-columns)

This works in most cases, but there are times when we just want to allow individual 0 values to be updated, and neither `map[string]interface` nor `Select` is very friendly to us.

# Usage

1. `NewPlugin` register plugin to `gorm.DB`

   ```go
       db.Use(zerofield.NewPlugin())
   ```

2. `UpdateScopes` update event it's zero field

   ```go

       // ...
       user.Name = ""
       user.Age = 0
       user.Active = false
       user.Birthday = nil

       // will always update Name,Age even if it's zero field
       // Active,Birthday will not be saved
       db.Scopes(zerofield.UpdateScopes("Name","Age")).Updates(&user)
       // if cloumns is empty, all field will be save like db.Select("*"")
       db.Scopes(zerofield.UpdateScopes()).Updates(&user)
   ```

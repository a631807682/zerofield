# zerofield

zero field pulgin for gorm.

# Desc

> When update with struct, GORM will only update non-zero fields, you might want to use map to update attributes or use Select to specify fields to update
> [Updates-multiple-columns](https://gorm.io/docs/update.html#Updates-multiple-columns)

This works in most cases, but there are times when we just want to allow individual 0 values to be updated, and neither `map[string]interface` nor `Select` is very friendly to us.

# Usage

## Scopes

1. `UpdateScopes` update event it's zero field

   ```go
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

## Plugins

1. `NewPlugin` force save by gorm tag `zf_force_update`

   > This is a dangerous option, usually `UpdateScopes` is enough

   ```go
       type Foo struct{
           Bar string `gorm:zf_force_update;`// will always update even if it's zero field
       }

       db.Use(zerofield.NewPlugin())

       foo.Bar = ""
       db.Updates(&foo)
   ```

# zerofield(WIP)

zero field pulgin for gorm.

# Desc

> When update with struct, GORM will only update non-zero fields, you might want to use map to update attributes or use Select to specify fields to update
> [Updates-multiple-columns](https://gorm.io/docs/update.html#Updates-multiple-columns)

This works in most cases, but there are times when we just want to allow individual 0 values to be updated, and neither `map[string]interface` nor `Select` is very friendly to us.

# Usage

## Scopes

1. `UpdateZeroFields` update event it's zero field

```go
    user.Name = ""
    user.Age = 0
    user.Active = false
    user.Birthday = nil

    // will always update Name,Age even if it's zero field
    // Active,Birthday will not be saved
    db.Scopes(scopes.UpdateZeroFields("Name","Age")).Updates(&user)
    // if cloumns is empty, all field will be save like db.Select("*"")
    db.Scopes(scopes.UpdateZeroFields()).Updates(&user)
```

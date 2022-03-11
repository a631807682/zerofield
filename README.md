# zerofield(WIP)

zero field pulgin for gorm.

# Usage

1. Update zero field by `Scopes`

```go
db.Scopes(zerofield.UpdateScope(zerofield.Config{})).Updates(&user)
```

2. Update zero field by `Hooks`

```go
db.Use(zerofield.UpdateHooks(zerofield.Config{}))
```

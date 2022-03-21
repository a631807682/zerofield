package utils

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Age       uint `gorm:"default:18"`
	Birthday  *time.Time
	Account   *Account
	Active    bool
	NotUpdate string `gorm:"<-:create"`
}

type Account struct {
	gorm.Model
	UserID sql.NullInt64
	Number string
}

type Foo struct {
	gorm.Model
	NotEmpty               string
	RestIfLonggerThan1Char string
	SwitchOn               bool `gorm:"zf_force_update:true"`
	SwitchStatus           int8 `gorm:"zf_force_update"`
}

const FooNotEmptyDefVal = "not empty"

func (f *Foo) BeforeUpdate(_ *gorm.DB) (err error) {
	if f.NotEmpty == "" {
		f.NotEmpty = FooNotEmptyDefVal
	}

	if len(f.RestIfLonggerThan1Char) > 1 {
		f.RestIfLonggerThan1Char = ""
	}
	return
}

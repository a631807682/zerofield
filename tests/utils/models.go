package utils

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      uint `gorm:"default:18"`
	Birthday *time.Time
	Account  *Account
	Pets     []*Pet
	Active   bool
}

type Account struct {
	gorm.Model
	UserID sql.NullInt64
	Number string
}

type Pet struct {
	gorm.Model
	UserID *uint
	Name   string
}

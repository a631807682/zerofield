module github.com/a631807682/zerofield/tests

go 1.17

replace github.com/a631807682/zerofield => ../

require (
	github.com/a631807682/zerofield v0.0.0-00010101000000-000000000000
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gorm v1.24.6
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
)

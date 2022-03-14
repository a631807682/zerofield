package tests_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
	. "zerofield/tests/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	var err error
	if DB, err = OpenTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	} else {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Printf("failed to connect database, got error %v", err)
			os.Exit(1)
		}

		err = sqlDB.Ping()
		if err != nil {
			log.Printf("failed to ping sqlDB, got error %v", err)
			os.Exit(1)
		}

		err = RunMigrations()
		if err != nil {
			log.Printf("failed to RunMigrations, got error %v", err)
			os.Exit(1)
		}
		if DB.Dialector.Name() == "sqlite" {
			DB.Exec("PRAGMA foreign_keys = ON")
		}
	}
}

// There is no need to test other databases because there is no driver interaction
func OpenTestConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	if err != nil {
		return
	}

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else if debug == "false" {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}

	return
}

func RunMigrations() error {
	var err error
	allModels := []interface{}{&User{}, &Account{}}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })

	if err = DB.Migrator().DropTable(allModels...); err != nil {
		return fmt.Errorf("Failed to drop table, got error %v\n", err)
	}

	if err = DB.AutoMigrate(allModels...); err != nil {
		return fmt.Errorf("Failed to auto migrate, but got error %v\n", err)
	}

	for _, m := range allModels {
		if !DB.Migrator().HasTable(m) {
			return fmt.Errorf("Failed to create table for %#v\n", m)
		}
	}
	return nil
}

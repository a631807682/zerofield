package utils

import (
	"regexp"
	"strings"
	"testing"

	"gorm.io/gorm"
)

// AssertEqualSQL for assert that the sql is equal, this method will ignore quote, and dialect speicals.
func AssertEqualSQL(t *testing.T, db *gorm.DB, expected string, actually string) {
	t.Helper()

	// replace SQL quote, convert into postgresql like ""
	expected = replaceQuoteInSQL(db, expected)
	actually = replaceQuoteInSQL(db, actually)

	// ignore updated_at value, becase it's generated in Gorm inernal, can't to mock value on update.
	updatedAtRe := regexp.MustCompile(`(?i)"updated_at"=".+?"`)
	actually = updatedAtRe.ReplaceAllString(actually, `"updated_at"=?`)
	expected = updatedAtRe.ReplaceAllString(expected, `"updated_at"=?`)

	// ignore RETURNING "id" (only in PostgreSQL)
	returningRe := regexp.MustCompile(`(?i)RETURNING "id"`)
	actually = returningRe.ReplaceAllString(actually, ``)
	expected = returningRe.ReplaceAllString(expected, ``)

	actually = strings.TrimSpace(actually)
	expected = strings.TrimSpace(expected)

	if actually != expected {
		t.Fatalf("\nexpected: %s\nactually: %s", expected, actually)
	}
}

func replaceQuoteInSQL(db *gorm.DB, sql string) string {
	// convert single quote into double quote
	sql = strings.ReplaceAll(sql, `'`, `"`)

	// convert dialect speical quote into double quote
	switch db.Dialector.Name() {
	case "postgres":
		sql = strings.ReplaceAll(sql, `"`, `"`)
	case "mysql", "sqlite":
		sql = strings.ReplaceAll(sql, "`", `"`)
	case "sqlserver":
		sql = strings.ReplaceAll(sql, `'`, `"`)
	}

	return sql
}

package tests_test

import (
	"testing"

	"github.com/a631807682/zerofield"
	. "github.com/a631807682/zerofield/tests/utils"

	gut "gorm.io/gorm/utils/tests"
)

func TestZeroFieldPluginTag(t *testing.T) {
	DB.Use(zerofield.NewPlugin())

	foo := Foo{
		RestIfLonggerThan1Char: "1",
		NotEmpty:               "notempty",
		SwitchOn:               true,
		SwitchStatus:           2,
	}

	DB.Create(&foo)

	foo.RestIfLonggerThan1Char = "2"
	foo.SwitchOn = false
	foo.SwitchStatus = 0
	DB.Updates(&foo)

	var foo1 Foo
	DB.First(&foo1, foo.ID)
	gut.AssertEqual(t, &foo1, &foo)
}

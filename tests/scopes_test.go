package tests_test

import (
	"testing"
	"time"

	"zerofield"

	. "zerofield/tests/utils"

	"gorm.io/gorm"
	gut "gorm.io/gorm/utils/tests"
)

func TestUpdateScopes(t *testing.T) {
	birthday := time.Now()
	user := User{
		Name:     "TestUpdateScopes",
		Age:      10,
		Active:   true,
		Birthday: &birthday,
	}
	DB.Create(&user)
	user.Name = ""
	user.Age = 0
	user.Active = false
	user.Birthday = nil

	DB.Scopes(zerofield.UpdateScopes()).Updates(&user)

	var user1 User
	DB.First(&user1, user.ID)
	gut.AssertEqual(t, &user1, &user)
}

func TestUpdateScopesWithField(t *testing.T) {
	birthday := time.Now()
	user := User{
		Name:     "TestUpdateScopesWithIncludes",
		Age:      10,
		Active:   true,
		Birthday: &birthday,
	}
	DB.Create(&user)
	user.Name = ""
	user.Age = 0
	user.Active = false
	user.Birthday = nil

	DB.Scopes(zerofield.UpdateScopes("Name")).Updates(&user)

	var user1 User
	DB.First(&user1, user.ID)
	gut.AssertObjEqual(t, &user1, &user, "ID", "Name")
	gut.AssertEqual(t, &user1.Age, 10)
	gut.AssertEqual(t, &user1.Active, true)
	gut.AssertEqual(t, &user1.Birthday, &birthday)
}

func TestUpdateScopesWithSelect(t *testing.T) {
	birthday := time.Now()
	user := User{
		Name:     "TestUpdateScopesWithSelect",
		Age:      10,
		Active:   true,
		Birthday: &birthday,
	}
	DB.Create(&user)
	user.Name = ""
	user.Age = 0
	user.Active = false
	user.Birthday = nil
	// update Active only
	DB.Scopes(zerofield.UpdateScopes()).Select("Active").Updates(&user)

	var user1 User
	DB.First(&user1, user.ID)
	gut.AssertEqual(t, &user1.Active, &user.Active)
}

func TestUpdateScopesWithOmit(t *testing.T) {
	birthday := time.Now()
	user := User{
		Name:     "TestUpdateScopesWithOmit",
		Age:      10,
		Active:   true,
		Birthday: &birthday,
	}
	DB.Create(&user)
	user.Name = ""
	user.Age = 0
	user.Active = false
	user.Birthday = nil

	// dont update Active only
	DB.Scopes(zerofield.UpdateScopes()).Omit("Active").Updates(&user)

	var user1 User
	DB.First(&user1, user.ID)
	gut.AssertEqual(t, &user1.Name, &user.Name)
	gut.AssertEqual(t, &user1.Age, &user.Age)
	gut.AssertEqual(t, &user1.Birthday, &user.Birthday)
	gut.AssertEqual(t, &user1.Active, true)
}

func TestUpdateScopesWithSpecifiedField(t *testing.T) {
	birthday := time.Now()
	user := User{
		Name:     "TestUpdateScopesWithInterface",
		Age:      10,
		Active:   true,
		Birthday: &birthday,
	}
	DB.Create(&user)

	sess := DB.Model(&user).Scopes(zerofield.UpdateScopes()).Session(&gorm.Session{})

	// donot handle
	sess.Update("name", "")
	var user1 User
	DB.First(&user1, user.ID)
	gut.AssertEqual(t, &user1.Name, "")

	// donot handle
	sess.Model(&user).Updates(map[string]interface{}{"age": 0, "active": false})
	var user2 User
	DB.First(&user2, user.ID)
	gut.AssertEqual(t, &user2.Age, 0)
	gut.AssertEqual(t, &user2.Active, false)
}

func TestUpdateScopesWithBeforeUpdateHooks(t *testing.T) {
	foo := &Foo{
		RestIfLonggerThan1Char: "1",
		NotEmpty:               "notempty",
	}

	DB.Create(&foo)

	foo.RestIfLonggerThan1Char = "morethanone"
	foo.NotEmpty = ""
	DB.Scopes(zerofield.UpdateScopes("RestIfLonggerThan1Char")).Updates(&foo)

	var foo1 Foo
	DB.First(&foo1, foo.ID)
	gut.AssertEqual(t, &foo1.RestIfLonggerThan1Char, "")
	gut.AssertEqual(t, &foo1.NotEmpty, FooNotEmptyDefVal)
}

func TestUpdateScopesWithSession(t *testing.T) {
	DB = DB.Debug()
	birthday := time.Now()
	user := User{
		Name:     "TestUpdateScopesWithInterface",
		Age:      10,
		Active:   true,
		Birthday: &birthday,
	}
	DB.Create(&user)

	sess := DB.Model(&user).Scopes(zerofield.UpdateScopes("Name", "Age")).Session(&gorm.Session{})

	user.Name = ""
	sess.Updates(&user)

	var user1 User
	DB.First(&user1, user.ID)
	gut.AssertEqual(t, &user1, &user)

	user.Age = 0
	sess.Updates(&user)

	var user2 User
	DB.First(&user2, user.ID)
	gut.AssertEqual(t, &user2, &user)
}

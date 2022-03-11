package tests_test

import (
	"testing"
	"time"

	"zerofield/scopes"

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

	DB.Scopes(scopes.UpdateScopes()).Updates(&user)

	var user1 User
	DB.First(&user1, user.ID)
	gut.AssertEqual(t, &user1, &user)
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
	DB.Scopes(scopes.UpdateScopes()).Select("Active").Updates(&user)

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
	DB.Scopes(scopes.UpdateScopes()).Omit("Active").Updates(&user)

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

	sess := DB.Model(&user).Scopes(scopes.UpdateScopes()).Session(&gorm.Session{})

	sess.Update("name", "")
	var user1 User
	DB.First(&user1, user.ID)
	gut.AssertEqual(t, &user1.Name, "")

	sess.Model(&user).Updates(map[string]interface{}{"age": 0, "active": false})
	var user2 User
	DB.First(&user2, user.ID)
	gut.AssertEqual(t, &user2.Age, 0)
	gut.AssertEqual(t, &user2.Active, false)
}

func TestUpdateScopesWithAssociations(t *testing.T) {
	user := User{
		Name: "TestUpdateScopesWithAssociations",
		Account: &Account{
			Number: "TestUpdateScopesWithAssociations_account",
		},
		Pets: []*Pet{
			{Name: "TestUpdateScopesWithAssociations_pet1"},
			{Name: "TestUpdateScopesWithAssociations_pet2"},
		},
	}
	DB.Create(&user)
	user.Account.Number = ""
	user.Account = nil
	user.Pets = user.Pets[1:]
	user.Age = 0
	// DB.Select(clause.Associations).Save(&user)
	DB.Scopes(scopes.UpdateScopes(&scopes.Config{
		Associations: true,
	})).Save(&user)
	// DB.Select("Account", "Pets").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

}

package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/verdade/golang-api-user/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateNewUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Ricardo", "ric@ig", "123123")
	userDb := NewUser(db)

	err = userDb.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
}

func TestUser_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Ricardo", "ric@ig", "123123")
	userDb := NewUser(db)

	err = userDb.Create(user)
	assert.Nil(t, err)

	userFound, err := userDb.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.NotEmpty(t, userFound.Name)

}

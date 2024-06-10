package database

import (
	"testing"

	"github.com/masilvasql/pos-go-expert-2024/APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John", "jonh@teste.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User

	err = db.Where("id = ?", user.ID).First(&userFound).Error

	assert.Nil(t, err)
	assert.Equal(t, "John", userFound.Name)
	assert.Equal(t, "jonh@teste.com", userFound.Email)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Password, user.Password)

}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John", "jonh@teste.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail("jonh@teste.com")

	assert.Nil(t, err)
	assert.Equal(t, "John", userFound.Name)
	assert.Equal(t, "jonh@teste.com", userFound.Email)
	assert.Equal(t, userFound.ID, user.ID)
	assert.NotNil(t, userFound.Password)
}

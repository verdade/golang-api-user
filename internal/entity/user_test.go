package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Ricardo", "ric@ig", "123123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Ricardo", user.Name)
	assert.Equal(t, "ric@ig", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Ricardo", "ric@ig", "123123")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123123"))
	assert.False(t, user.ValidatePassword("34"))
	assert.NotEqual(t, "123123", user.Password)
}

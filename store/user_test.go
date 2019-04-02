package store

import (
	"galileo/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserStore_FindByEmail(t *testing.T) {
	s := NewUserStore()
	u := types.NewUser()
	u.Email = "user@email.com"
	s.Add(u)

	u1, err := s.FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.Equal(t, u, u1)
}

func TestUserStore_FindByToken(t *testing.T) {
	s := NewUserStore()
	u := types.NewUser()
	s.Add(u)

	u1, err := s.FindByToken(u.Token)
	assert.NoError(t, err)
	assert.Equal(t, u, u1)
}

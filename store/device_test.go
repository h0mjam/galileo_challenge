package store

import (
	"galileo/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeviceStore_FindByID(t *testing.T) {
	s := NewDeviceStore()
	d := types.NewDevice()

	s.Add(d)

	d1, err := s.FindByID(d.ID)

	assert.NoError(t, err)
	assert.Equal(t, d, d1)
}

func TestDeviceStore_FindByToken(t *testing.T) {
	s := NewDeviceStore()
	d := types.NewDevice()

	s.Add(d)

	d1, err := s.FindByToken(d.Token)

	assert.NoError(t, err)
	assert.Equal(t, d, d1)
}

package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDevice_AppendMeasure(t *testing.T) {
	d := NewDevice()

	tt, err := time.Parse("", "")

	assert.NoError(t, err)

	m := &Measure{
		Value:     1.01,
		CreatedAt: tt,
	}

	d.AppendMeasure(m)

	m1, ok := d.Measures[tt.UTC().Unix()]

	assert.Equal(t, ok, true)
	assert.Equal(t, m, m1)
}

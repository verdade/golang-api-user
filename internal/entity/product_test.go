package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("P1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "P1", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10.0)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("23", 0.0)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
}

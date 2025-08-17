package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validDC, _ = NewDistributionCenter("e28b5578-41c6-4911-81f1-e4da8efb6aba", "CD12", "S2", "SP", "ACTIVE", 100)

func TestCreateANewProductWithSuccess(t *testing.T) {
	p, err := NewProduct("ea16c354-abe1-416f-9f28-68dc4a0b5c0d", "Prod Name", 12.23, 3, *validDC)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, "ea16c354-abe1-416f-9f28-68dc4a0b5c0d", p.CatalogProductId)
	assert.Equal(t, "Prod Name", p.Name)
	assert.Equal(t, 12.23, p.Price)
	assert.Equal(t, 3, p.Quantity)
}
func TestCreateANewProductWithEmptyName(t *testing.T) {
	p, err := NewProduct("ea16c354-abe1-416f-9f28-68dc4a0b5c0d", "", 12.23, 3, *validDC)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrProductInvalidName, err)
}
func TestCreateANewProductWithInvalidPrice(t *testing.T) {
	p, err := NewProduct("ea16c354-abe1-416f-9f28-68dc4a0b5c0d", "Prod Name", -4.23, 3, *validDC)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrProductInvalidPrice, err)
}

func TestCreateANewProductWithInvalidQuantity(t *testing.T) {
	p, err := NewProduct("ea16c354-abe1-416f-9f28-68dc4a0b5c0d", "Prod Name", 2.3, -23, *validDC)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrProductInvalidQuantity, err)
}

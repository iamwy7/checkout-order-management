package domain

import (
	"testing"

	"github.com/iamwy7/meli-challenge/orders/application/shared"
	"github.com/stretchr/testify/assert"
)

var validProducts = *orderProducts(1)

func TestCreateANewOrderWithSuccess(t *testing.T) {
	o, err := NewOrder(validProducts, "S2", "SP")
	assert.NotNil(t, o)
	assert.Nil(t, err)
	assert.Equal(t, "ea16c354-abe1-416f-9f28-68dc4a0b5c0d", o.Products[0].CatalogProductId)
	assert.Equal(t, shared.South2, o.Zone)
	assert.Equal(t, "SP", o.State)
}

func TestCreateANewOrderWithNilProducts(t *testing.T) {
	var nilProducts []Product
	o, err := NewOrder(nilProducts, "S2", "SP")
	assert.NotNil(t, err)
	assert.Nil(t, o)
	assert.Equal(t, ErrOrderWithoutProducts, err)
}
func TestCreateANewOrderWithAHundredPlusProducts(t *testing.T) {
	o, err := NewOrder(*orderProducts(105), "S2", "SP")
	assert.NotNil(t, err)
	assert.Nil(t, o)
	assert.Equal(t, ErrOrderWithTooMuchProducts, err)
}
func TestCreateANewOrderWithInvalidState(t *testing.T) {
	o, err := NewOrder(validProducts, "S2", "BA")
	assert.NotNil(t, err)
	assert.Nil(t, o)
	assert.Equal(t, ErrOrderInvalidState, err)
}

func TestCreateANewOrderWithInvalidZone(t *testing.T) {
	o, err := NewOrder(validProducts, "INVALID_ZONE", "SP")
	assert.NotNil(t, err)
	assert.Nil(t, o)
	assert.Equal(t, ErrOrderInvalidZone, err)
}

func orderProducts(howMuch int) *[]Product {
	var hundredProds = make([]Product, howMuch)
	for i := range howMuch {
		p, _ := NewProduct("ea16c354-abe1-416f-9f28-68dc4a0b5c0d", "Prod Name", 12.23, 3, *validDC)
		hundredProds[i] = *p
	}
	return &hundredProds
}

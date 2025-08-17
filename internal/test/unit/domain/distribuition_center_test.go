package domain

import (
	"testing"

	"github.com/iamwy7/meli-challenge/orders/application/shared"
	"github.com/stretchr/testify/assert"
)

func TestCreateANewDistributionCenterWithSuccess(t *testing.T) {
	dc, err := NewDistributionCenter("e28b5578-41c6-4911-81f1-e4da8efb6aba", "CD12", "S2", "SP", "ACTIVE", 100)
	assert.Nil(t, err)
	assert.NotNil(t, dc)
	assert.Equal(t, "e28b5578-41c6-4911-81f1-e4da8efb6aba", dc.Id)
	assert.Equal(t, "CD12", dc.Name)
	assert.Equal(t, shared.South2, dc.Zone)
	assert.Equal(t, "SP", dc.State)
	assert.Equal(t, shared.ACTIVE, dc.Status)
	assert.Equal(t, 100, dc.ProductQuantity)
}

func TestCreateANewDistributionCenterWithEmptyId(t *testing.T) {
	// test nil or empty  id
	dc, err := NewDistributionCenter("", "CD12", "S2", "SP", "ACTIVE", 100)
	assert.NotNil(t, err)
	assert.Nil(t, dc)
	assert.Equal(t, ErrDCInvalidId, err)
}
func TestCreateANewDistributionCenterWithEmptyName(t *testing.T) {
	// test nil or empty a name
	dc, err := NewDistributionCenter("e28b5578-41c6-4911-81f1-e4da8efb6aba", "", "S2", "SP", "ACTIVE", 100)
	assert.Nil(t, dc)
	assert.NotNil(t, err)
	assert.Equal(t, ErrDCNameRequired, err)
}
func TestCreateANewDistributionCenterWithInvalidZone(t *testing.T) {
	// test with invalid zone
	dc, err := NewDistributionCenter("e28b5578-41c6-4911-81f1-e4da8efb6aba", "CD12", "INVALID_ZONE", "SP", "ACTIVE", 100)
	assert.Nil(t, dc)
	assert.NotNil(t, err)
	assert.Equal(t, ErrDCInvalidZone, err)
}
func TestCreateANewDistributionCenterWithEmptyZone(t *testing.T) {
	// test nil or empty zone
	dc, err := NewDistributionCenter("e28b5578-41c6-4911-81f1-e4da8efb6aba", "CD12", "", "SP", "ACTIVE", 100)
	assert.Nil(t, dc)
	assert.NotNil(t, err)
	assert.Equal(t, ErrDCInvalidZone, err)
}
func TestCreateANewDistributionCenterWithInvalidStatus(t *testing.T) {
	// test with invalid status
	dc, err := NewDistributionCenter("e28b5578-41c6-4911-81f1-e4da8efb6aba", "CD12", "S2", "SP", "INVALID_STATUS", 100)
	assert.Nil(t, dc)
	assert.NotNil(t, err)
	assert.Equal(t, ErrDcInvalidStatus, err)
}
func TestCreateANewDistributionCenterWithEmptyStatus(t *testing.T) {
	// test nil or empty status
	dc, err := NewDistributionCenter("e28b5578-41c6-4911-81f1-e4da8efb6aba", "CD12", "S2", "SP", "", 100)
	assert.Nil(t, dc)
	assert.NotNil(t, err)
	assert.Equal(t, ErrDcInvalidStatus, err)
}
func TestCreateANewDistributionCenterWithInvalidProductQuantity(t *testing.T) {
	// test invalid product quantity
	dc, err := NewDistributionCenter("e28b5578-41c6-4911-81f1-e4da8efb6aba", "CD12", "S2", "SP", "ACTIVE", -23)
	assert.Nil(t, dc)
	assert.NotNil(t, err)
	assert.Equal(t, ErrDCInvalidQuantity, err)
}

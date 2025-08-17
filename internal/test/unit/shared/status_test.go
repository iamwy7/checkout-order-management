package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAValidStatus(t *testing.T) {
	zone := CheckStatus("ACTIVE")
	assert.Equal(t, ACTIVE, zone)
}

func TestInValidStatuses(t *testing.T) {
	t.Run("CheckInvalidStatusS123", func(t *testing.T) {
		zone := CheckStatus("S123")
		assert.Equal(t, "", string(zone))
	})
	t.Run("CheckInvalidStatusINVALID_ZONE", func(t *testing.T) {
		zone := CheckStatus("INVALID_ZONE")
		assert.Equal(t, "", string(zone))
	})
	t.Run("CheckEmptyStatus", func(t *testing.T) {
		zone := CheckStatus("")
		assert.Equal(t, "", string(zone))
	})
}

package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Zone is the zone where the client is located.
func TestAValidZone(t *testing.T) {
	zone := CheckZone("S1")
	assert.Equal(t, South1, zone)
}

func TestInValidZones(t *testing.T) {
	t.Run("CheckInvalidZoneS123", func(t *testing.T) {
		zone := CheckZone("S123")
		assert.Equal(t, "", string(zone))
	})
	t.Run("CheckInvalidZoneINVALID_ZONE", func(t *testing.T) {
		zone := CheckZone("INVALID_ZONE")
		assert.Equal(t, "", string(zone))
	})
	t.Run("CheckEmptyZone", func(t *testing.T) {
		zone := CheckZone("")
		assert.Equal(t, "", string(zone))
	})
}

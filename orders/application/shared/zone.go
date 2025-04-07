package shared

// Zone is the zone where the client is located.
type Zone string

// List of zones available on SP.
const (
	North1 Zone = "N1"
	North2 Zone = "N2"
	East1  Zone = "E1"
	East2  Zone = "E2"
	West   Zone = "W1"
	Center Zone = "C1"
	South1 Zone = "S1"
	South2 Zone = "S2"
)

func ValidateZone(zone string) (Zone, error) {
	if zone != string(North1) && zone != string(North2) && zone != string(East1) &&
		zone != string(East2) && zone != string(West) && zone != string(Center) &&
		zone != string(South1) && zone != string(South2) {
		return "", ErrInvalidZone
	}
	return Zone(zone), nil
}

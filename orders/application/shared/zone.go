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

func CheckZone(zone string) Zone {
	switch Zone(zone) {
	case North1, North2, East1, East2, West, Center, South1, South2:
		return Zone(zone)
	default:
		return ""
	}
}

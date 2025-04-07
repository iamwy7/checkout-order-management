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

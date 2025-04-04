package domain

type DistributionCenter struct {
	Id   string // UUID
	Name string // CD1, CD2 etc...
	Zone Zone
}

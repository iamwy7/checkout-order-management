package shared

type Status string

const (
	ACTIVE   Status = "ACTIVE"
	INACTIVE Status = "APPROVED"
	DELETED  Status = "DELETED"
	PENDING  Status = "PENDING"
	APPROVED Status = "APPROVED"
	CANCELED Status = "CANCELED"
)

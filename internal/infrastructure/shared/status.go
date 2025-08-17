package shared

// Status are the available statuses to order and distribution center operation status
type Status string

const (
	ACTIVE   Status = "ACTIVE"
	INACTIVE Status = "INACTIVE"
	DELETED  Status = "DELETED"
	PENDING  Status = "PENDING"
	APPROVED Status = "APPROVED"
	CANCELED Status = "CANCELED"
)

func CheckStatus(status string) Status {
	switch Status(status) {
	case ACTIVE, INACTIVE, DELETED, PENDING, APPROVED, CANCELED:
		return Status(status)
	default:
		return ""
	}
}

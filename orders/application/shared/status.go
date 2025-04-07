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

// ValidateStatus checks if the status is valid.
func ValidateStatus(status string) (Status, error) {
	if status != string(ACTIVE) && status != string(INACTIVE) && status != string(DELETED) &&
		status != string(PENDING) && status != string(APPROVED) && status != string(CANCELED) {
		return "", ErrInvalidStatus
	}
	return Status(status), nil
}

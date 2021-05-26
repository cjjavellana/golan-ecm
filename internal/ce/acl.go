package ce

type AccessControlType string

const (
	// AccessControlTypeGroup represents a group ownership. Anybody who is a member of a group
	// whose group name appears in the `Owner` field of the AccessControl struct can perform
	// the actions as per described in the PermissionType
	AccessControlTypeGroup AccessControlType = "group"
	AccessControlTypeUser                    = "user"
)

type AccessControl struct {
	Owner         string
	OwnershipType AccessControlType
	Permission    PermissionType
}

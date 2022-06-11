package interfaces

type InterfaceMethod string

const (
	CollectionsQuery  InterfaceMethod = "CollectionsQuery"
	CollectionsWrite  InterfaceMethod = "CollectionsWrite"
	CollectionsCommit InterfaceMethod = "CollectionsCommit"
	CollectionsDelete InterfaceMethod = "CollectionsDelete"

	PermissionsGrant InterfaceMethod = "PermissionsGrant"
)

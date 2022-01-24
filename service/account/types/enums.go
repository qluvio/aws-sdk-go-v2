// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AlternateContactType string

// Enum values for AlternateContactType
const (
	AlternateContactTypeBilling    AlternateContactType = "BILLING"
	AlternateContactTypeOperations AlternateContactType = "OPERATIONS"
	AlternateContactTypeSecurity   AlternateContactType = "SECURITY"
)

// Values returns all known values for AlternateContactType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AlternateContactType) Values() []AlternateContactType {
	return []AlternateContactType{
		"BILLING",
		"OPERATIONS",
		"SECURITY",
	}
}
package models

type PathPermissionsWithCount struct {
	PathPermissions PathPermissions `json:"pathPermissions"`
	Count           int             `json:"count"`
}

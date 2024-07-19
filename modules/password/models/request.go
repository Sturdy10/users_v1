package models

type InitPassword struct {
	OrgmbEmail  string `json:"email"`
	NewPassword string `json:"newPassword"`
}

type ChangePassword struct {
	OrgmbEmail      string `json:"email"`
	Oldpassword     string `json:"oldPassword"`
	Newpassword     string `json:"newPassword"`

}
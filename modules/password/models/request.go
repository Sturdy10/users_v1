package models

type InitPassword struct {
	MdUsername  string `json:"username"`
	NewPassword string `json:"newPassword"`
}

type ChangePassword struct {
	MdUsername      string `json:"username"`
	Oldpassword     string `json:"oldPassword"`
	Newpassword     string `json:"newPassword"`

}
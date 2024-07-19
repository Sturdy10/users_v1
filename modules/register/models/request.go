package models

type RegisterMember struct {
	OrgmbTitle        string `json:"title"`
	OrgmbName         string `json:"name"`
	OrgmbSurname      string `json:"surname"`
	OrgmbEmail        string `json:"email"`
	OrgmbMobile       string `json:"mobile"`
	OrgdpName         string `json:"department"`
	GeneratedPassword string
}

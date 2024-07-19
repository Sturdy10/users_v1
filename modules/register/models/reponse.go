package models

type MemberResponse struct {
	OrgmbID      string `json:"member_id"`
	OrgmbTitle   string `json:"title"`
	OrgmbName    string `json:"name"`
	OrgmbSurname string `json:"surname"`
	OrgmbEmail   string `json:"email"`
	OrgmbMobile  string `json:"mobile"`
	OrgrlOrgdpID    string `json:"role_orgdp_id"`
	OrgdpName    string `json:"department"`
}

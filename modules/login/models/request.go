package models


type LoginRequest struct {
	MbUsername string `json:"username"`
	MbPassword string `json:"password"`
}

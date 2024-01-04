package dto

type LoginResponse struct {
	Token string `json:"token"`
	//ExpiredAt string `json:"expired"`
}

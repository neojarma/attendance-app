package model

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

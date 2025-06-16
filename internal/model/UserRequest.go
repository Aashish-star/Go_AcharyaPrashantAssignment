package model

type UserRequest struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

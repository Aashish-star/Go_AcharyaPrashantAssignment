package model

type UserResponse struct {
	IsSucess     bool   `json: "isSuccess"`
	Message      string `json: "message"`
	AccessToken  string `json: "accessToken"`
	RefreshToken string `json: "refreshToken"`
}

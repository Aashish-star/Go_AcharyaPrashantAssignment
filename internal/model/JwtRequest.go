package model

type JwtRequest struct {
	RefreshToken string `json: "refreshToken"`
}

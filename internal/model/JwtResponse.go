package model

type JwtResponse struct {
	AccessToken  string `json: "accessToken"`
	RefreshToken string `json: "refreshToken"`
}

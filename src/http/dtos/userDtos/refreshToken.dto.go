package userDtos

type RefreshTokenDto struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

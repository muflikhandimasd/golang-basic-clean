package userDtos

import (
	userModels "dmp-training/models"
)

type LoginDto struct {
	Id           int32  `json:"id"`
	Username     string `json:"username"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	LastLoginAt  string `json:"last_login_at"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func NewLoginDto(v *userModels.User, token, refreshToken string) *LoginDto {
	return &LoginDto{
		Id:           v.Id,
		Username:     v.Username,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
		LastLoginAt:  v.LastLoginAt,
		Token:        token,
		RefreshToken: refreshToken,
	}

}

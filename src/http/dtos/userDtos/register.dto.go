package userDtos

type RegisterDto LoginDto

func NewRegisterDto(v *LoginDto) *RegisterDto {
	return &RegisterDto{
		Id:           v.Id,
		Username:     v.Username,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
		LastLoginAt:  v.LastLoginAt,
		Token:        v.Token,
		RefreshToken: v.RefreshToken,
	}
}

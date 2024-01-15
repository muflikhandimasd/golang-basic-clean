package models

import "database/sql"

const ColumnsUser = "id,username,password,created_at,updated_at,last_login_at"

type User struct {
	Id          int32  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	LastLoginAt string `json:"last_login_at"`
}

type UserSql struct {
	Id          sql.NullInt32  `json:"id"`
	Username    sql.NullString `json:"username"`
	Password    sql.NullString `json:"password"`
	CreatedAt   sql.NullString `json:"created_at"`
	UpdatedAt   sql.NullString `json:"updated_at"`
	LastLoginAt sql.NullString `json:"last_login_at"`
}

func (u *UserSql) ToUser() *User {
	return &User{
		Id:          u.Id.Int32,
		Username:    u.Username.String,
		Password:    u.Password.String,
		CreatedAt:   u.CreatedAt.String,
		UpdatedAt:   u.UpdatedAt.String,
		LastLoginAt: u.LastLoginAt.String,
	}

}

func (u *UserSql) Scanners() []interface{} {
	return []interface{}{
		&u.Id,
		&u.Username,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.LastLoginAt,
	}
}

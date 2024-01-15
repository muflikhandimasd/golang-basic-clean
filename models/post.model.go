package models

import "database/sql"

const ColumnsPost = "id, title, body, created_at, updated_at,user_id"

type Post struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserId    int32  `json:"user_id"`
}

type PostSql struct {
	ID        sql.NullInt32  `json:"id"`
	Title     sql.NullString `json:"title"`
	Body      sql.NullString `json:"body"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
	UserId    sql.NullInt32  `json:"user_id"`
}

func (p *PostSql) Scanners() []interface{} {
	return []interface{}{
		&p.ID,
		&p.Title,
		&p.Body,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.UserId,
	}
}

func (p *PostSql) ToPost() *Post {
	return &Post{
		ID:        p.ID.Int32,
		Title:     p.Title.String,
		Body:      p.Body.String,
		CreatedAt: p.CreatedAt.String,
		UpdatedAt: p.UpdatedAt.String,
		UserId:    p.UserId.Int32,
	}
}

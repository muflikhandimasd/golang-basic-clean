package mysql

import (
	"context"
	"database/sql"

	"github.com/muflikhandimasd/golang-basic-clean/configs"
	"github.com/muflikhandimasd/golang-basic-clean/database"
	"github.com/muflikhandimasd/golang-basic-clean/models"
	"github.com/sirupsen/logrus"
)

type PostRepository interface {
	GetAll(ctx context.Context, userId int32) (response []*models.Post, err error)
	Create(ctx context.Context, title, body string, userId int32) (err error)
	Update(ctx context.Context, id int, title, body string) (err error)
	Delete(ctx context.Context, id int) (err error)
	IsExists(ctx context.Context, id int) (exist bool, err error)
	IsExistsByTitle(ctx context.Context, title string) (exist bool, err error)
}

type postRepository struct {
	db  *sql.DB
	log *logrus.Entry
	tag string
}

func NewPostRepository(cfg configs.Config) PostRepository {
	return &postRepository{
		db:  cfg.DB(),
		log: cfg.Log(),
		tag: "ERROR DB - POST REPOSITORY",
	}
}

func (r *postRepository) GetAll(ctx context.Context, userId int32) (response []*models.Post, err error) {
	const TAG = "GetAll"

	sqlString := "SELECT " + models.ColumnsPost + " FROM posts WHERE user_id = ?"

	err = database.WithTransaction(ctx, r.db, func(tx database.Transaction) error {

		rows, err := tx.QueryContext(ctx, sqlString, userId)
		if err != nil {
			r.log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error(r.tag + " - " + TAG)

			return err
		}

		defer rows.Close()

		for rows.Next() {
			post := new(models.PostSql)
			err = rows.Scan(
				post.Scanners()...,
			)

			if err != nil {
				r.log.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Error(r.tag + " - " + TAG)

				return err
			}

			response = append(response, post.ToPost())
		}

		return nil
	})

	return
}

func (r *postRepository) Create(ctx context.Context, title, body string, userId int32) (err error) {
	const TAG = "Create"

	sqlString := "INSERT INTO posts (title,body,user_id) VALUES (?,?,?)"

	args := []interface{}{
		title,
		body,
		userId,
	}

	err = database.WithTransaction(ctx, r.db, func(tx database.Transaction) error {
		_, err := tx.ExecContext(ctx, sqlString, args...)
		if err != nil {
			r.log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error(r.tag + " - " + TAG)

			return err
		}

		return nil
	})

	return
}

func (r *postRepository) Update(ctx context.Context, id int, title, body string) (err error) {
	const TAG = "Update"

	sqlString := "UPDATE posts SET title = ?, body = ? WHERE id = ?"

	args := []interface{}{
		title,
		body,
		id,
	}

	err = database.WithTransaction(ctx, r.db, func(tx database.Transaction) error {
		_, err := tx.ExecContext(ctx, sqlString, args...)
		if err != nil {
			r.log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error(r.tag + " - " + TAG)

			return err
		}

		return nil
	})

	return
}

func (r *postRepository) Delete(ctx context.Context, id int) (err error) {
	const TAG = "Delete"

	sqlString := "DELETE FROM posts WHERE id = ?"

	args := []interface{}{
		id,
	}

	err = database.WithTransaction(ctx, r.db, func(tx database.Transaction) error {
		_, err := tx.ExecContext(ctx, sqlString, args...)
		if err != nil {
			r.log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error(r.tag + " - " + TAG)

			return err
		}

		return nil
	})

	return
}

func (r *postRepository) IsExists(ctx context.Context, id int) (exist bool, err error) {
	const TAG = "IsExist"

	sqlString := "SELECT COUNT(id) FROM posts WHERE id = ?"

	args := []interface{}{
		id,
	}

	err = database.WithTransaction(ctx, r.db, func(tx database.Transaction) error {
		var total int

		err = tx.QueryRowContext(ctx, sqlString, args...).Scan(&total)
		if err != nil {
			r.log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error(r.tag + " - " + TAG)

			return err
		}

		if total > 0 {
			exist = true
		}

		return nil
	})

	return
}

func (r *postRepository) IsExistsByTitle(ctx context.Context, title string) (exist bool, err error) {
	const TAG = "IsExist"

	sqlString := "SELECT COUNT(title) FROM posts WHERE title = ?"

	args := []interface{}{
		title,
	}

	err = database.WithTransaction(ctx, r.db, func(tx database.Transaction) error {
		var total int

		err = tx.QueryRowContext(ctx, sqlString, args...).Scan(&total)
		if err != nil {
			r.log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error(r.tag + " - " + TAG)

			return err
		}

		if total > 0 {
			exist = true
		}
		return nil
	})

	return
}

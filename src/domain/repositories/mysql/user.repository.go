package mysql

import (
	"context"
	"database/sql"
	"dmp-training/configs"
	"dmp-training/database"
	"dmp-training/helpers"
	userModels "dmp-training/models"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int64) (response []*userModels.User, err error)
	GetUserByUsername(ctx context.Context, username string) (response []*userModels.User, err error)
	Register(ctx context.Context, username, password string) (id int64, err error)
	CreateRefreshToken(ctx context.Context, userId int32, refreshToken string) (err error)
}

type userRepository struct {
	db  *sql.DB
	log *logrus.Entry
	tag string
}

func NewUserRepository(cfg configs.Config) UserRepository {
	return &userRepository{
		db:  cfg.DB(),
		log: cfg.Log(),
		tag: "ERROR DB - USER REPOSITORY",
	}
}

func (r *userRepository) GetUserById(ctx context.Context, id int64) (response []*userModels.User, err error) {
	const TAG = "GetUserById"

	sqlString := "SELECT " + userModels.ColumnsUser + " FROM users WHERE id = ? FOR UPDATE"

	args := []interface{}{
		id,
	}

	err = database.WithTransaction(ctx, r.db, func(tx database.Transaction) error {

		rows, err := tx.QueryContext(ctx, sqlString, args...)
		if err != nil {
			r.log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error(r.tag + " - " + TAG)

			return err
		}

		defer rows.Close()

		for rows.Next() {
			user := new(userModels.UserSql)
			err = rows.Scan(
				user.Scanners()...,
			)

			if err != nil {
				r.log.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Error(r.tag + " - " + TAG)

				return err
			}

			response = append(response, user.ToUser())
		}

		return nil
	})
	return
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (response []*userModels.User, err error) {
	const TAG = "GetUserByUsername"

	sqlString := "SELECT " + userModels.ColumnsUser + " FROM users WHERE username = ? FOR UPDATE"

	args := []interface{}{
		username,
	}

	err = database.WithTransaction(ctx, r.db, func(tx database.Transaction) error {

		rows, err := tx.QueryContext(ctx, sqlString, args...)
		if err != nil {
			r.log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error(r.tag + " - " + TAG)

			return err
		}

		defer rows.Close()

		for rows.Next() {
			user := new(userModels.UserSql)
			err = rows.Scan(
				user.Scanners()...,
			)

			if err != nil {
				r.log.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Error(r.tag + " - " + TAG)

				return err
			}

			response = append(response, user.ToUser())
		}

		return nil
	})
	return
}

func (r *userRepository) Register(ctx context.Context, username, password string) (id int64, err error) {
	const TAG = "Register"

	sqlString := "INSERT INTO users (username,password) VALUES (?,?)"

	args := []interface{}{
		username,
		helpers.GenerateHash(password),
	}

	err = database.WithTransaction(ctx, r.db, func(tx database.Transaction) error {

		result, err := tx.ExecContext(ctx, sqlString, args...)
		if err != nil {
			r.log.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error(r.tag + " - " + TAG)

			return err
		}

		id, err = result.LastInsertId()

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

func (r *userRepository) CreateRefreshToken(ctx context.Context, userId int32, refreshToken string) (err error) {
	const TAG = "CreateRefreshToken"

	sqlString := "INSERT INTO refresh_tokens (user_id,refresh_token) VALUES (?,?) ON DUPLICATE KEY UPDATE refresh_token = ? "

	args := []interface{}{
		userId,
		refreshToken,
		refreshToken,
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

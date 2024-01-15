package usecases

import (
	"context"
	"dmp-training/configs"
	"dmp-training/constants"
	"dmp-training/helpers"
	"dmp-training/src/domain/repositories/mysql"
	"dmp-training/src/http/dtos/userDtos"
	userRequests "dmp-training/src/http/requests/userRequests"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

type UserUseCase interface {
	Login(ctx context.Context, request *userRequests.LoginRequest) (code int, msg string, res *userDtos.LoginDto)
	Register(ctx context.Context, request *userRequests.RegisterRequest) (code int, msg string, res *userDtos.RegisterDto)
	RefreshToken(ctx context.Context, request *userRequests.RefreshTokenRequest) (code int, msg string, res *userDtos.RefreshTokenDto)
}

type userUseCase struct {
	userRepo   mysql.UserRepository
	log        *logrus.Entry
	ctxTimeout time.Duration
	tag        string
}

func NewUserUseCase(userRepo mysql.UserRepository, config configs.Config) UserUseCase {
	return &userUseCase{
		userRepo:   userRepo,
		log:        config.Log(),
		ctxTimeout: config.ContextTimeout(),
		tag:        "ERROR USECASE - UserUseCase",
	}
}

func (r *userUseCase) Login(ctx context.Context, request *userRequests.LoginRequest) (code int, msg string, res *userDtos.LoginDto) {
	time.Local = time.FixedZone("UTC+7", 7*60*60)

	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)

	defer cancel()

	code = fiber.StatusOK
	msg = constants.MessageSuccess

	users, err := r.userRepo.GetUserByUsername(ctx, request.Username)

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - Login")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	if len(users) == 0 {
		code = fiber.StatusUnauthorized
		msg = constants.MessageUnauthorized
		return
	}

	user := users[0]

	if ok := helpers.CompareHash(request.Password, user.Password); !ok {
		code = fiber.StatusUnauthorized
		msg = constants.MessageUnauthorized
		return
	}

	token := helpers.GenerateToken(user.Id, user.Username)

	refreshToken := helpers.GenerateRefreshToken(user.Id, user.Username)

	err = r.userRepo.CreateRefreshToken(ctx, user.Id, refreshToken)

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - RefreshToken")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	res = userDtos.NewLoginDto(user, token, refreshToken)
	return
}

func (r *userUseCase) Register(ctx context.Context, request *userRequests.RegisterRequest) (code int, msg string, res *userDtos.RegisterDto) {
	time.Local = time.FixedZone("UTC+7", 7*60*60)

	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)

	defer cancel()

	code = fiber.StatusOK
	msg = constants.MessageSuccess

	users, err := r.userRepo.GetUserByUsername(ctx, request.Username)

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - Register")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	if len(users) > 0 {
		code = fiber.StatusConflict
		msg = constants.MessageConflict
		return
	}

	id, err := r.userRepo.Register(ctx, request.Username, request.Password)

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - Register")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	users, err = r.userRepo.GetUserById(ctx, id)

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - Register")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	user := users[0]

	token := helpers.GenerateToken(user.Id, user.Username)

	refreshToken := helpers.GenerateRefreshToken(user.Id, user.Username)

	err = r.userRepo.CreateRefreshToken(ctx, user.Id, refreshToken)

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - RefreshToken")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	res = userDtos.NewRegisterDto(userDtos.NewLoginDto(user, token, refreshToken))
	return
}

func (r *userUseCase) RefreshToken(ctx context.Context, request *userRequests.RefreshTokenRequest) (code int, msg string, res *userDtos.RefreshTokenDto) {
	time.Local = time.FixedZone("UTC+7", 7*60*60)

	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)

	defer cancel()

	code = fiber.StatusOK
	msg = constants.MessageSuccess

	claim, err := helpers.ParseRefreshToken(request.RefreshToken)

	if err != nil {
		code = fiber.StatusUnauthorized
		msg = constants.MessageUnauthorized
		return
	}

	users, err := r.userRepo.GetUserById(ctx, int64(claim.Id))

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - RefreshToken")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	if len(users) == 0 {
		code = fiber.StatusUnauthorized
		msg = constants.MessageUnauthorized
		return
	}

	user := users[0]

	token := helpers.GenerateToken(user.Id, user.Username)

	refreshToken := helpers.GenerateRefreshToken(user.Id, user.Username)

	err = r.userRepo.CreateRefreshToken(ctx, user.Id, refreshToken)

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - RefreshToken")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	res = new(userDtos.RefreshTokenDto)
	res.Token = token
	res.RefreshToken = refreshToken
	return
}

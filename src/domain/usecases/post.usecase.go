package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/muflikhandimasd/golang-basic-clean/configs"
	"github.com/muflikhandimasd/golang-basic-clean/constants"
	"github.com/muflikhandimasd/golang-basic-clean/models"
	"github.com/muflikhandimasd/golang-basic-clean/src/domain/repositories/mysql"
	"github.com/muflikhandimasd/golang-basic-clean/src/http/requests/postRequests"
	"github.com/sirupsen/logrus"
)

type PostUseCase interface {
	GetAll(ctx context.Context, request *postRequests.GetAllPostRequest) (code int, msg string, res []*models.Post)
	Create(ctx context.Context, request *postRequests.CreatePostRequest) (code int, msg string)
	Update(ctx context.Context, request *postRequests.UpdatePostRequest) (code int, msg string)
	Delete(ctx context.Context, request *postRequests.DeletePostRequest) (code int, msg string)
}

type postUseCase struct {
	postRepo   mysql.PostRepository
	log        *logrus.Entry
	ctxTimeout time.Duration
	tag        string
}

func NewPostUseCase(postRepo mysql.PostRepository, cfg configs.Config) PostUseCase {
	return &postUseCase{
		postRepo:   postRepo,
		log:        cfg.Log(),
		ctxTimeout: cfg.ContextTimeout(),
		tag:        "ERROR USECASE - PostUseCase",
	}
}

func (r *postUseCase) GetAll(ctx context.Context, request *postRequests.GetAllPostRequest) (code int, msg string, res []*models.Post) {
	time.Local = time.FixedZone("UTC+7", 7*60*60)

	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)

	defer cancel()

	code = fiber.StatusOK
	msg = constants.MessageSuccess

	res, err := r.postRepo.GetAll(ctx, request.UserId)
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - GetAll")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	if len(res) == 0 {
		res = make([]*models.Post, 0)
		return
	}

	return
}

func (r *postUseCase) Create(ctx context.Context, request *postRequests.CreatePostRequest) (code int, msg string) {
	time.Local = time.FixedZone("UTC+7", 7*60*60)

	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)

	defer cancel()

	code = fiber.StatusOK
	msg = constants.MessageSuccess

	exists, err := r.postRepo.IsExistsByTitle(ctx, request.Title)
	fmt.Println("exists", exists)
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - Create")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	if exists {
		code = fiber.StatusConflict
		msg = constants.MessageDataAlreadyExists
		return
	}

	err = r.postRepo.Create(ctx, request.Title, request.Body, request.UserId)
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - Create")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	return
}

func (r *postUseCase) Update(ctx context.Context, request *postRequests.UpdatePostRequest) (code int, msg string) {
	time.Local = time.FixedZone("UTC+7", 7*60*60)

	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)

	defer cancel()

	code = fiber.StatusOK
	msg = constants.MessageSuccess

	exists, err := r.postRepo.IsExists(ctx, request.Id)

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - Update")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	if !exists {
		code = fiber.StatusNotFound
		msg = constants.MessageDataNotFound
		return
	}

	exists, err = r.postRepo.IsExistsByTitle(ctx, request.Title)

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - Update")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	if exists {
		code = fiber.StatusConflict
		msg = constants.MessageDataAlreadyExists
		return
	}

	err = r.postRepo.Update(ctx, request.Id, request.Title, request.Body)
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - Update")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	return
}

func (r *postUseCase) Delete(ctx context.Context, request *postRequests.DeletePostRequest) (code int, msg string) {
	time.Local = time.FixedZone("UTC+7", 7*60*60)

	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)

	defer cancel()

	code = fiber.StatusOK
	msg = constants.MessageSuccess

	exists, err := r.postRepo.IsExists(ctx, request.Id)

	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - Update")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	if !exists {
		code = fiber.StatusNotFound
		msg = constants.MessageDataNotFound
		return
	}

	err = r.postRepo.Delete(ctx, request.Id)
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(r.tag + " - Update")
		code = fiber.StatusInternalServerError
		msg = constants.MessageInternalServerError
		return
	}

	return
}

package src

import (
	"dmp-training/configs"
	"dmp-training/registries"
	"dmp-training/src/http/controllers"
	"dmp-training/src/http/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

type server struct {
	http *fiber.App
	cfg  configs.Config
	uc   registries.UsecaseRegistry
}

type Server interface {
	Run()
}

func NewServer(cfg configs.Config) Server {
	app := fiber.New(
		fiber.Config{
			EnablePrintRoutes: true,
		},
	)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	app.Use(middlewares.NewLogger())
	app.Use(middlewares.NewLimiter())
	app.Use(middlewares.NewIP())

	repo := registries.NewRepositoryRegistry(cfg)

	uc := registries.NewUsecaseRegistry(repo, cfg)

	return &server{
		http: app,
		cfg:  cfg,
		uc:   uc,
	}

}

func (s *server) Run() {
	apiGroup := s.http.Group("/api")

	userController := controllers.NewAuthController(s.uc.User())

	userController.Groups(apiGroup)

	postController := controllers.NewPostController(s.uc.Post())

	postController.Groups(apiGroup)

	log.Fatal(s.http.Listen(":" + s.cfg.Port()))
}

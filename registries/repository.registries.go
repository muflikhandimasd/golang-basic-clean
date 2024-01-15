package registries

import (
	"dmp-training/configs"
	"dmp-training/src/domain/repositories/mysql"
	"sync"
)

type repositoryRegistry struct {
	cfg configs.Config
}

type RepositoryRegistry interface {
	User() mysql.UserRepository
	Post() mysql.PostRepository
}

func NewRepositoryRegistry(cfg configs.Config) RepositoryRegistry {
	return &repositoryRegistry{
		cfg: cfg,
	}
}

func (r *repositoryRegistry) User() mysql.UserRepository {
	var repo mysql.UserRepository
	var loadonce sync.Once

	loadonce.Do(func() {
		repo = mysql.NewUserRepository(r.cfg)
	})

	return repo
}

func (r *repositoryRegistry) Post() mysql.PostRepository {
	var repo mysql.PostRepository
	var loadonce sync.Once

	loadonce.Do(func() {
		repo = mysql.NewPostRepository(r.cfg)
	})

	return repo
}

package registries

import (
	"dmp-training/configs"
	"dmp-training/src/domain/usecases"
	"sync"
)

type usecaseRegistry struct {
	repo RepositoryRegistry
	cfg  configs.Config
}

type UsecaseRegistry interface {
	User() usecases.UserUseCase
	Post() usecases.PostUseCase
}

func NewUsecaseRegistry(repo RepositoryRegistry, cfg configs.Config) UsecaseRegistry {
	return &usecaseRegistry{
		repo: repo,
		cfg:  cfg,
	}
}

func (r *usecaseRegistry) User() usecases.UserUseCase {
	var uc usecases.UserUseCase
	var loadonce sync.Once

	loadonce.Do(func() {
		uc = usecases.NewUserUseCase(r.repo.User(), r.cfg)
	})

	return uc
}

func (r *usecaseRegistry) Post() usecases.PostUseCase {
	var uc usecases.PostUseCase
	var loadonce sync.Once

	loadonce.Do(func() {
		uc = usecases.NewPostUseCase(r.repo.Post(), r.cfg)
	})

	return uc
}

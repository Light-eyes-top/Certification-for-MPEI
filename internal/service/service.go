package service

import (
	service_models "certification/internal/models/service-models"
	"certification/internal/repository"
)

type User interface {
	CreateUser(user *service_models.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Sign interface {
	CreateSign(buffer []byte, userId int) error
	CheckSign(buffer []byte, userId int) (bool, error)
	DeleteSign(buffer []byte, userId int) error
}

type Service struct {
	User
	Sign
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo.User),
		Sign: NewSignService(repo.Sign),
	}
}

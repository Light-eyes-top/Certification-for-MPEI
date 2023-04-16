package repository

import (
	db_models "certification/intrenal/models/db-models"
	service_models "certification/intrenal/models/service-models"
	"certification/intrenal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user *db_models.User) (int, error)
	GetUser(username, password string) (*service_models.User, error)
}

type Sign interface {
	CreateSign(hash string, userId int) error
	CheckSign(hash string, userId int) (bool, error)
	DeleteSign(hash string, userId int) error
}

type Repository struct {
	Sign
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: postgres.NewUserRepository(db),
		Sign: postgres.NewSignRepository(db),
	}
}

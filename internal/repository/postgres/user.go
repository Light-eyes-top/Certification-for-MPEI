package postgres

import (
	db_models "certification/internal/models/db-models"
	"certification/internal/models/mapper"
	service_models "certification/internal/models/service-models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(user *db_models.User) (int, error) {
	var id int
	line := u.db.QueryRow("INSERT INTO users (name, password_hash) VALUES ($1, $2) RETURNING id",
		user.Username, user.Password)
	if err := line.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserRepository) GetUser(username, password string) (*service_models.User, error) {
	var user db_models.User
	line := fmt.Sprintf("SELECT * FROM users WHERE name = $1 AND password_hash = $2")
	err := u.db.Get(&user, line, username, password)

	return mapper.UserDbToService(&user), err
}

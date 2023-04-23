package postgres

import (
	"github.com/jmoiron/sqlx"
)

type SignRepository struct {
	db *sqlx.DB
}

func NewSignRepository(db *sqlx.DB) *SignRepository {
	return &SignRepository{db: db}
}

func (u *SignRepository) CreateSign(hash string, userId int) error {
	query := "INSERT INTO signs (hash, user_id) VALUES ($1, $2)"
	_, err := u.db.Exec(query, hash, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u *SignRepository) CheckSign(hash string, userId int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT * FROM signs s WHERE s.hash = $1 AND s.user_id = $2)"
	err := u.db.Get(&exists, query, hash, userId)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *SignRepository) DeleteSign(hash string, userId int) error {
	query := "DELETE FROM signs s WHERE s.hash = $1 AND s.user_id = $2"
	_, err := u.db.Exec(query, hash, userId)
	return err
}

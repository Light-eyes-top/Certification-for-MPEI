package db_models

type User struct {
	Id       int    `db:"id"`
	Username string `db:"name"`
	Password string `db:"password_hash"`
}

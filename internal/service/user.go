package service

import (
	"certification/internal/consts"
	"certification/internal/models/mapper"
	service_models "certification/internal/models/service-models"
	"certification/internal/repository"
	"crypto/sha1"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserService struct {
	user repository.User
}

func NewUserService(user repository.User) *UserService {
	return &UserService{user: user}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (u *UserService) CreateUser(user *service_models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return u.user.CreateUser(mapper.UserServiceToDb(user))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hash.Sum([]byte(consts.Salt)))

	return sha
}

func (u *UserService) GenerateToken(username, password string) (string, error) {
	user, err := u.user.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(consts.TokenTLL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(consts.SigningKey))
}

package service

import (
	"certification/internal/repository"
	"crypto/sha256"
	"encoding/base64"
)

type SignService struct {
	sign repository.Sign
}

func NewSignService(sign repository.Sign) *SignService {
	return &SignService{sign: sign}
}

func (s *SignService) CreateSign(buffer []byte, userId int) error {
	buffer = append(buffer, byte(userId))
	byteHash := sha256.Sum256(buffer)
	hash := base64.URLEncoding.EncodeToString(byteHash[:])
	return s.sign.CreateSign(hash, userId)
}

func (s *SignService) CheckSign(buffer []byte, userId int) (bool, error) {
	buffer = append(buffer, byte(userId))
	byteHash := sha256.Sum256(buffer)
	hash := base64.URLEncoding.EncodeToString(byteHash[:])
	return s.sign.CheckSign(hash, userId)
}

func (s *SignService) DeleteSign(buffer []byte, userId int) error {
	buffer = append(buffer, byte(userId))
	byteHash := sha256.Sum256(buffer)
	hash := base64.URLEncoding.EncodeToString(byteHash[:])
	return s.sign.DeleteSign(hash, userId)
}

func (s *SignService) UserSigned() {

}

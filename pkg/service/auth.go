package service

import (
	"crypto/sha1"
	"fmt"

	rsapi "github.com/wanna-beat-by-bit/rest-golang"
	"github.com/wanna-beat-by-bit/rest-golang/pkg/repository"
)

const salt = "dihtkladsfh123osfdio"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user rsapi.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

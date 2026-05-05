package service

import (
	"finance-api/internal/model"
	"finance-api/internal/repository"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(name, email, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userInfo := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}

	return s.userRepo.CreateUser(userInfo)
}

func (s *AuthService) Login(email, password string) (string, error) {
	emailInfo, _ := s.userRepo.GetUserByEmail(email)
	if emailInfo == nil {
		return "", fmt.Errorf("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(emailInfo.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": emailInfo.ID,
		"exp":     jwt.TimeFunc().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

package service

import (
	"errors"

	"github.com/harshdevops117/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type LoginService struct {
	db *gorm.DB
}

func NewLoginService(db *gorm.DB) *LoginService {
	return &LoginService{db: db}
}

func (s *LoginService) Login(email, password string) (*models.User, error) {
	var user models.User

	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return nil, ErrInvalidCredentials
	}
	return &user, nil
}

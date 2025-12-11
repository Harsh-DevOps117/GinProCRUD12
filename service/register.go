package service

import (
	"errors"

	"github.com/harshdevops117/dto"
	"github.com/harshdevops117/models"
	"github.com/harshdevops117/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterService struct {
	db *gorm.DB
}

func NewRegisterService(db *gorm.DB) *RegisterService {
	return &RegisterService{db: db}
}

func (s *RegisterService) Register(body dto.RegisterDTO) (*models.User, error) {

	if !validator.EmailValidator(body.Email) {
		return nil, errors.New("invalid email")
	}

	if !validator.PasswordValidator(body.Password) {
		return nil, errors.New("invalid password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

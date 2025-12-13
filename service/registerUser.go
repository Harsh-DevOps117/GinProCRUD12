package service

import (
	"errors"

	"github.com/harshdevops117/models"
	"github.com/harshdevops117/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUser struct {
	db *gorm.DB
}
func NewRegisterUser(db *gorm.DB) *RegisterUser {
	return &RegisterUser{db: db}
}

func (r *RegisterUser) RegisterUser(data *models.User) error {
	ValidEmail:=validator.EmailValidator(data.Email)
	if(!ValidEmail){
		return nil
	}
	hashPassword,err:=bcrypt.GenerateFromPassword([]byte(data.Password),10)
	if err!=nil {
		 return errors.New("invalid email")
	}
	data.Password=string(hashPassword)
	newUser:=models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	if err := r.db.Create(&newUser).Error; err != nil {
		return err
	}
	return nil
}

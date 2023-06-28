package repository

import (
	"errors"

	"github.com/rafialariq/go-bank/model"
	"github.com/rafialariq/go-bank/model/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRepository interface {
	GetUser(*dto.LoginDTO) (*model.User, error)
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) LoginRepository {
	return &loginRepository{db}
}

func (l *loginRepository) GetUser(user *dto.LoginDTO) (*model.User, error) {
	var existUser model.User

	err := l.db.Where("phone_number = ?", user.PhoneNumber).First(&existUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &existUser, errors.New("user not found")
		}
		// logging here
		return &existUser, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password))
	if err != nil {
		// logging here
		return &existUser, errors.New("username or password is not valid")
	}

	return &existUser, nil
}

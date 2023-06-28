package repository

import (
	"errors"
	"fmt"

	"github.com/rafialariq/go-bank/model"
	"github.com/rafialariq/go-bank/model/dto"
	"gorm.io/gorm"
)

type RegisterRepository interface {
	InsertUser(*model.User) error
	FindExistingUser(*dto.RegisterDTO) bool
}

type registerRepository struct {
	db *gorm.DB
}

func NewRegisterRepository(db *gorm.DB) RegisterRepository {
	return &registerRepository{db}
}

func (r *registerRepository) InsertUser(user *model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		fmt.Println("error 1", err) // temporary
		return err
	}

	return nil
}

func (r *registerRepository) FindExistingUser(user *dto.RegisterDTO) bool {
	var existUser model.User
	err := r.db.Where("username = ? OR phone_number = ?", user.Username, user.PhoneNumber).First(&existUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("error 2", err) // temporary
			return true
		}

		// logging here
		// not completed
		fmt.Println("error 3", err) // temporray
	}

	return false
}

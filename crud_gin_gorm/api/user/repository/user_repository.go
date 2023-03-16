package repository

import (
	"crud_gin_gorm/api/user/model"
	"crud_gin_gorm/database/config"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	config.InitDB()
	return &UserRepository{db: config.GetDB()}
}

func (u *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) FindByID(userID int) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Save(user *model.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Update(userID int, user *model.User) error {
	if err := u.db.Model(&model.User{}).Where("id = ?", userID).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Delete(userID int) error {
	if err := u.db.Where("id = ?", userID).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}

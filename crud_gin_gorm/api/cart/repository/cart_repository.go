package repository

import (
	"crud_gin_gorm/api/cart/model"
	"crud_gin_gorm/database/config"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository() *CartRepository {
	config.InitDB()
	return &CartRepository{db: config.GetDB()}
}

func (u *CartRepository) FindAll() ([]model.Cart, error) {
	var carts []model.Cart
	if err := u.db.Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (u *CartRepository) FindByID(cartID int) (*model.Cart, error) {
	var cart model.Cart
	if err := u.db.First(&cart, cartID).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (u *CartRepository) FindByUserID(userID int) (*model.Cart, error) {
	var cart model.Cart
	if err := u.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (u *CartRepository) Save(cart *model.Cart) error {
	if err := u.db.Create(&cart).Error; err != nil {
		return err
	}
	return nil
}

func (u *CartRepository) Update(cartID int, cart *model.Cart) error {
	if err := u.db.Model(&model.Cart{}).Where("id = ?", cartID).Updates(&cart).Error; err != nil {
		return err
	}
	return nil
}

func (u *CartRepository) Delete(cartID int) error {
	if err := u.db.Where("id = ?", cartID).Delete(&model.Cart{}).Error; err != nil {
		return err
	}
	return nil
}

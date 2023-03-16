package repository

import (
	"crud_gin_gorm/api/cart/model"
	"crud_gin_gorm/database/config"

	"gorm.io/gorm"
)

type CartItemRepository struct {
	db *gorm.DB
}

func NewCartItemRepository() *CartItemRepository {
	config.InitDB()
	return &CartItemRepository{db: config.GetDB()}
}

func (u *CartItemRepository) FindAll() ([]model.CartItem, error) {
	var cartItems []model.CartItem
	if err := u.db.Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (u *CartItemRepository) FindByID(cartItemID int) (*model.CartItem, error) {
	var cartItem model.CartItem
	if err := u.db.First(&cartItem, cartItemID).Error; err != nil {
		return nil, err
	}
	return &cartItem, nil
}

func (u *CartItemRepository) FindByCartID(cartID int) ([]*model.CartItem, error) {
	var cartItems []*model.CartItem
	if err := u.db.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (u *CartItemRepository) Save(cartItem *model.CartItem) error {
	if err := u.db.Create(&cartItem).Error; err != nil {
		return err
	}
	return nil
}

func (u *CartItemRepository) Update(cartItemID int, cartItem *model.CartItem) error {
	if err := u.db.Model(&model.CartItem{}).Where("id = ?", cartItemID).Updates(&cartItem).Error; err != nil {
		return err
	}
	return nil
}

func (u *CartItemRepository) Delete(cartItemID int) error {
	if err := u.db.Where("id = ?", cartItemID).Delete(&model.CartItem{}).Error; err != nil {
		return err
	}
	return nil
}

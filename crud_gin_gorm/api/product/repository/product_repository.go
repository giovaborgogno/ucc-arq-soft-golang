package repository

import (
	"crud_gin_gorm/api/product/model"
	"crud_gin_gorm/database/config"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() *ProductRepository {
	config.InitDB()
	return &ProductRepository{db: config.GetDB()}
}

func (u *ProductRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	if err := u.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (u *ProductRepository) FindByID(productID int) (*model.Product, error) {
	var product model.Product
	if err := u.db.First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (u *ProductRepository) Save(product *model.Product) error {
	if err := u.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (u *ProductRepository) Update(productID int, product *model.Product) error {
	if err := u.db.Model(&model.Product{}).Where("id = ?", productID).Updates(&product).Error; err != nil {
		return err
	}
	return nil
}

func (u *ProductRepository) Delete(productID int) error {
	if err := u.db.Where("id = ?", productID).Delete(&model.Product{}).Error; err != nil {
		return err
	}
	return nil
}

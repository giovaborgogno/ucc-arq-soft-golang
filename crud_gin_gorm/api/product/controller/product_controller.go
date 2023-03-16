package controller

import (
	"net/http"
	"strconv"

	"crud_gin_gorm/api/common"
	"crud_gin_gorm/api/product/model"
	"crud_gin_gorm/api/product/repository"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

var (
	productRepo = repository.NewProductRepository()
)

func (u *ProductController) GetProducts(c *gin.Context) {
	products, err := productRepo.FindAll()
	if err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error retrieving products")
		return
	}
	common.JSONResponse(c, http.StatusOK, products)
}

func (u *ProductController) GetProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}
	product, err := productRepo.FindByID(productID)
	if err != nil {
		common.ErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}
	common.JSONResponse(c, http.StatusOK, product)
}

func (u *ProductController) CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := productRepo.Save(&product); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error saving product")
		return
	}
	common.JSONResponse(c, http.StatusCreated, product)
}

func (u *ProductController) UpdateProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := productRepo.Update(productID, &product); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error updating product")
		return
	}
	productIDnew := uint(productID)
	product.ID = productIDnew
	common.JSONResponse(c, http.StatusOK, product)
}

func (u *ProductController) DeleteProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}
	if err := productRepo.Delete(productID); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error deleting product")
		return
	}
	common.JSONResponse(c, http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

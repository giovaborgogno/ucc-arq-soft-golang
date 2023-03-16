package controller

import (
	"net/http"
	"strconv"

	"crud_gin_gorm/api/cart/model"
	"crud_gin_gorm/api/cart/repository"
	"crud_gin_gorm/api/common"
	productrepo "crud_gin_gorm/api/product/repository"

	"github.com/gin-gonic/gin"
)

type CartController struct{}

var (
	cartRepo     = repository.NewCartRepository()
	cartItemRepo = repository.NewCartItemRepository()
	productRepo  = productrepo.NewProductRepository()
)

func (u *CartController) GetCarts(c *gin.Context) {
	carts, err := cartRepo.FindAll()
	if err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error retrieving carts")
		return
	}
	common.JSONResponse(c, http.StatusOK, carts)
}

func (u *CartController) GetCart(c *gin.Context) {
	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid cart ID")
		return
	}
	cart, err := cartRepo.FindByUserID(cartID)
	if err != nil {
		common.ErrorResponse(c, http.StatusNotFound, "Cart not found")
		return
	}
	common.JSONResponse(c, http.StatusOK, cart)
}

func (u *CartController) CreateCart(c *gin.Context) {
	var cart model.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := cartRepo.Save(&cart); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error saving cart")
		return
	}
	common.JSONResponse(c, http.StatusCreated, cart)
}

func (u *CartController) UpdateCart(c *gin.Context) {
	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}
	var cart model.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := cartRepo.Update(cartID, &cart); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error updating cart")
		return
	}
	common.JSONResponse(c, http.StatusOK, cart)
}

func (u *CartController) DeleteCart(c *gin.Context) {
	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid cart ID")
		return
	}
	if err := cartRepo.Delete(cartID); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error deleting cart")
		return
	}
	common.JSONResponse(c, http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}

func (u *CartController) GetCartItems(c *gin.Context) {
	cartItems, err := cartItemRepo.FindAll()
	if err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error retrieving cart items")
		return
	}
	common.JSONResponse(c, http.StatusOK, cartItems)
}

func (u *CartController) GetCartItem(c *gin.Context) {
	cartItemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid cart item ID")
		return
	}
	cartItem, err := cartItemRepo.FindByID(cartItemID)
	if err != nil {
		common.ErrorResponse(c, http.StatusNotFound, "Cart item not found")
		return
	}
	common.JSONResponse(c, http.StatusOK, cartItem)
}

func (u *CartController) GetCartItemsByCart(c *gin.Context) {
	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid cart ID")
		return
	}
	cartItems, err := cartItemRepo.FindByCartID(cartID)
	if err != nil {
		common.ErrorResponse(c, http.StatusNotFound, "Cart items not found")
		return
	}
	common.JSONResponse(c, http.StatusOK, cartItems)
}

func (u *CartController) CreateCartItem(c *gin.Context) {
	var cartItem model.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	product, err := productRepo.FindByID(int(cartItem.ProductID))
	if err != nil {
		common.ErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}
	cartItem.Price = product.Price * float64(cartItem.Quantity)
	if err := cartItemRepo.Save(&cartItem); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error saving cart item")
		return
	}
	common.JSONResponse(c, http.StatusCreated, cartItem)
}

func (u *CartController) UpdateCartItem(c *gin.Context) {
	cartItemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid cart item ID")
		return
	}
	var cartItem model.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	product, err := productRepo.FindByID(int(cartItem.ProductID))
	if err != nil {
		common.ErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}
	cartItem.Price = product.Price * float64(cartItem.Quantity)
	uintCartItemID := uint(cartItemID)
	cartItem.ID = uintCartItemID
	if err := cartItemRepo.Update(cartItemID, &cartItem); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error updating cart item")
		return
	}
	common.JSONResponse(c, http.StatusOK, cartItem)
}

func (u *CartController) DeleteCartItem(c *gin.Context) {
	cartItemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid cart item ID")
		return
	}
	if err := cartItemRepo.Delete(cartItemID); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error deleting cart item")
		return
	}
	common.JSONResponse(c, http.StatusOK, gin.H{"message": "Cart item deleted successfully"})
}

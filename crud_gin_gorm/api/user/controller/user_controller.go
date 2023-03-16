package controller

import (
	"net/http"
	"strconv"

	cartmodel "crud_gin_gorm/api/cart/model"
	cartrepo "crud_gin_gorm/api/cart/repository"
	"crud_gin_gorm/api/common"
	"crud_gin_gorm/api/user/model"
	"crud_gin_gorm/api/user/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var (
	userRepo = repository.NewUserRepository()
	cartRepo = cartrepo.NewCartRepository()
)

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := userRepo.FindAll()
	if err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error retrieving users")
		return
	}
	common.JSONResponse(c, http.StatusOK, users)
}

func (u *UserController) GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user, err := userRepo.FindByID(userID)
	if err != nil {
		common.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	common.JSONResponse(c, http.StatusOK, user)
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := userRepo.Save(&user); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error saving user")
		return
	}
	var cart cartmodel.Cart
	cart.UserID = user.ID
	if err := cartRepo.Save(&cart); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error saving cart")
		return
	}
	common.JSONResponse(c, http.StatusCreated, user)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := userRepo.Update(userID, &user); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error updating user")
		return
	}
	userIDnew := uint(userID)
	user.ID = userIDnew
	common.JSONResponse(c, http.StatusOK, user)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}
	if err := userRepo.Delete(userID); err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, "Error deleting user")
		return
	}
	common.JSONResponse(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}

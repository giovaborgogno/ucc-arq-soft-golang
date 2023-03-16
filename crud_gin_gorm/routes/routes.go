package routes

import (
	cart "crud_gin_gorm/api/cart/controller"
	product "crud_gin_gorm/api/product/controller"
	user "crud_gin_gorm/api/user/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userController := user.UserController{}
	cartController := cart.CartController{}
	productController := product.ProductController{}

	v1 := r.Group("/api")
	{

		users := v1.Group("/users")
		{
			users.GET("", userController.GetUsers)
			users.POST("", userController.CreateUser)
			users.GET("/:id", userController.GetUser)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}

		products := v1.Group("/products")
		{
			products.GET("", productController.GetProducts)
			products.POST("", productController.CreateProduct)
			products.GET("/:id", productController.GetProduct)
			products.PUT("/:id", productController.UpdateProduct)
			products.DELETE("/:id", productController.DeleteProduct)
		}

		carts := v1.Group("/carts")
		{
			carts.GET("", cartController.GetCarts)
			carts.GET("/:id", cartController.GetCart)
			carts.GET("/items/:id", cartController.GetCartItemsByCart)
			carts.POST("/items", cartController.CreateCartItem)
			carts.PUT("/items/:id", cartController.UpdateCartItem)
			carts.DELETE("/items/:id", cartController.DeleteCartItem)

		}
	}

	return r
}

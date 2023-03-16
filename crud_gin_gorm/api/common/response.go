package common

import (
	// "net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse devuelve una respuesta de error con el mensaje y el código de estado HTTP proporcionados.
func ErrorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// JSONResponse devuelve una respuesta JSON con los datos y el código de estado HTTP proporcionados.
func JSONResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

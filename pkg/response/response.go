package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONResponse(ctx *gin.Context, statusCode int, success bool, message string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"Success": success,
		"Message": message,
		"Data":    data,
	})
}

func JSONErrorResponse(ctx *gin.Context, statusCode int, success bool, message string, err error) {
	ctx.JSON(statusCode, gin.H{
		"Success": success,
		"Message": message,
		"Error":   err.Error(),
	})
}

func JsonInputValidation(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"Success": false,
		"Message": "client-side input validation failed",
		"Error":   "Error in Binding the JSON Data",
	})
}

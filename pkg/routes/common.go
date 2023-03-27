package routes

import (
	"github.com/LucasCarioca/home-controls-services/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkKeyMiddleware(ctx *gin.Context) {
	config := config.GetConfig()
	apiKey := config.GetString("API_KEY")
	requestKey := ctx.Query("api_key")
	if apiKey != requestKey {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized request",
			"error":   "UNAUTHORIZED_REQUEST",
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}

package router

import (
	"main/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(walletHandler *handler.WalletHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/health", healthCheck)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/balance/:user_id", walletHandler.GetBalance)
		v1.POST("/withdraw", walletHandler.Withdraw)
	}

	return router
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Service is running",
	})
}
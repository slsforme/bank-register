package routers

import (
	"bankapp/internal/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

func registerTransactionRoutes(rg *gin.RouterGroup, h *handlers.Handler) {
	transactions := rg.Group("/transactions")
	{
		transactions.GET("", h.GetAllTransactions)
		transactions.GET("/user/:id", h.GetTransactionsByUserID)
		transactions.POST("", h.CreateTransaction)
	}
}

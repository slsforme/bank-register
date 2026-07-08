package routers

import (
	"bankapp/internal/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

func registerAccountEndpoints(rg *gin.RouterGroup, h *handlers.Handler) {
	accounts := rg.Group("/accounts")
	{
		accounts.POST("", h.CreateAccount)
		accounts.POST("/:id", h.UpdateAccount)
		accounts.POST("/:id", h.DeleteAccount)
		accounts.GET("/:id", h.GetAccountByID)
		accounts.GET("", h.GetAllAccounts)
		accounts.POST("", h.Transfer)
	}
}

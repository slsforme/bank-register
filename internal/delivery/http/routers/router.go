package routers

import (
	"bankapp/internal/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

func GetRouter(h *handlers.Handler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")

	registerTransactionRoutes(api, h)
	registerAccountEndpoints(api, h)

	return router
}

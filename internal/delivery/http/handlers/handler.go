package handlers

import (
	apperrors "bankapp/internal/delivery/errors"
	"bankapp/internal/usecase"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	accountUC  *usecase.AccountUseCase
	transferUC *usecase.TransferUseCase
	logger     *slog.Logger
}

func NewHandler(accountUC *usecase.AccountUseCase, transferUC *usecase.TransferUseCase, logger *slog.Logger) *Handler {
	return &Handler{
		accountUC:  accountUC,
		transferUC: transferUC,
		logger:     logger,
	}
}

func (h *Handler) handleError(c *gin.Context, err error) {
	switch err {

	case apperrors.InvalidOwnerName:
		h.logger.Info("invalid data in request",
			"error", apperrors.InvalidOwnerName,
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
			"ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
			"query", c.Request.URL.RawQuery,
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperrors.InvalidOwnerName.Error()})

	case apperrors.BalanceBelowZero:
		h.logger.Info("invalid data in request",
			"error", apperrors.BalanceBelowZero,
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
			"ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
			"query", c.Request.URL.RawQuery,
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": apperrors.BalanceBelowZero.Error()})

	default:
		h.logger.Error("internal error", "error",
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
			"ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
			"query", c.Request.URL.RawQuery,
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

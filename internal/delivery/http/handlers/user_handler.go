package handlers

import (
	apperrors "bankapp/internal/delivery/errors"
	"bankapp/internal/delivery/http/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len([]rune(req.OwnerName)) < 2 || len([]rune(req.OwnerName)) > 255 {
		h.handleError(c, apperrors.InvalidOwnerName)
		return
	}

	if req.Balance < 0 {
		h.handleError(c, apperrors.BalanceBelowZero)
		return
	}

	user, err := h.accountUC.CreateAccount(c.Request.Context(), req.OwnerName, req.Balance)

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) UpdateAccount(c *gin.Context) {
	var req dto.UpdateUserRequest

	idParam := c.Query("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		h.handleError(c, apperrors.InvalidID)
		return
	}

	if id == 0 {
		h.handleError(c, apperrors.InvalidID)
		return
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.OwnerName != nil && (len([]rune(*req.OwnerName)) < 2 || len([]rune(*req.OwnerName)) > 255) {
		h.handleError(c, apperrors.InvalidOwnerName)
		return
	}

	if req.Balance != nil && *req.Balance < 0 {
		h.handleError(c, apperrors.BalanceBelowZero)
		return
	}

	user, err := h.accountUC.CreateAccount(c.Request.Context(), *req.OwnerName, *req.Balance)

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteAccount(c *gin.Context) {
	var req dto.DeleteUserRequest

	idParam := c.Query("id")

	if idParam != "" {
		id, err := strconv.Atoi(idParam)

		if err != nil {
			h.handleError(c, apperrors.InvalidID)
			return
		}

		if id == 0 {
			h.handleError(c, apperrors.InvalidID)
			return
		}

		user, err := h.accountUC.DeleteAccount(c.Request.Context(), id)

		if err != nil {
			h.handleError(c, err)
			return
		}

		c.JSON(http.StatusCreated, user)
	} else {
		if err := c.ShouldBindBodyWithJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := h.accountUC.DeleteAccount(c.Request.Context(), req.ID)

		if err != nil {
			h.handleError(c, err)
			return
		}

		c.JSON(http.StatusOK, user)
	}

}

func (h *Handler) GetAccountByID(c *gin.Context) {
	var req dto.DeleteUserRequest

	idParam := c.Query("id")

	if idParam != "" {
		id, err := strconv.Atoi(idParam)

		if err != nil {
			h.handleError(c, apperrors.InvalidID)
			return
		}

		if id == 0 {
			h.handleError(c, apperrors.InvalidID)
			return
		}

		user, err := h.accountUC.GetAccount(c.Request.Context(), id)

		if err != nil {
			h.handleError(c, err)
			return
		}

		c.JSON(http.StatusCreated, user)
	} else {
		if err := c.ShouldBindBodyWithJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := h.accountUC.GetAccount(c.Request.Context(), req.ID)

		if err != nil {
			h.handleError(c, err)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func (h *Handler) GetAllAccounts(c *gin.Context) {
	var req dto.GetAllUsersRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := h.accountUC.GetAllAccounts(c.Request.Context())

	if req.Limit == 0 {
		req.Limit = 100
	}

	if req.Offset <= 0 {
		req.Offset = 0
	}

	paginatedUsers := (*users)[req.Offset:req.Limit]

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, paginatedUsers)
}

func (h *Handler) Transfer(c *gin.Context) {
	var req dto.TransferRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.SenderID == 0 || req.ReceiverID == 0 {
		h.handleError(c, apperrors.InvalidID)
		return
	}

	err := h.transferUC.Transfer(c, req.SenderID, req.ReceiverID, req.Amount)

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

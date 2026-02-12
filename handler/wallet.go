package handler

import (
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	service service.WalletService
}

func NewWalletHandler(service service.WalletService) *WalletHandler {
	return &WalletHandler{service: service}
}

type WithdrawRequest struct {
	UserID string  `json:"user_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type BalanceResponse struct {
	UserID  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	userID := c.Param("user_id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "user_id is required",
		})
		return
	}

	wallet, err := h.service.GetBalance(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, BalanceResponse{
		UserID:  wallet.UserID,
		Balance: wallet.Balance,
	})
}

func (h *WalletHandler) Withdraw(c *gin.Context) {
	var req WithdrawRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "withdrawal amount must be greater than zero",
		})
		return
	}

	err := h.service.Withdraw(req.UserID, req.Amount)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "insufficient balance" || err.Error() == "wallet not found" {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Withdrawal successful",
	})
}

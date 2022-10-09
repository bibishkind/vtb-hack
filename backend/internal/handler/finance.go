package handler

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetBalanceResponse struct {
	BalanceMatic  float32 `json:"balanceMatic"`
	BalanceRubles float32 `json:"balanceRubles"`
}

type TransferMaticRequest struct {
	ReceiverId int     `json:"receiverId"`
	Amount     float32 `json:"amount"`
}

type TransferRubleRequest struct {
	ReceiverId int     `json:"receiverId"`
	Amount     float32 `json:"amount"`
}

type TransferNftRequest struct {
	ReceiverId int
	TokenId    float32
}

// @Summary Get Balance
// @Security ApiKeyAuth
// @Tags finance
// @Description get balance of the user
// @Accept json
// @Produce json
// @Success 200 {object} GetBalanceResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/balance [get]
func (h *Handler) GetBalance(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	userId, ok := c.Get("userId").(int)
	if !ok {
		return makeErrorResponse(c, http.StatusUnauthorized, errors.New("invalid user id"))
	}

	balanceMatic, balanceRubles, err := h.service.GetBalance(ctx, userId)
	if err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"balanceMatic":  balanceMatic,
		"balanceRubles": balanceRubles,
	})
}

// @Summary Transfer Matic
// @Security ApiKeyAuth
// @Tags finance
// @Description transfer matic from current user to another
// @Accept json
// @Produce json
// @Param transferMatic body TransferMaticRequest true "transferMatic"
// @Success 200
// @Failure 401 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/transfer/matic [post]
func (h *Handler) TransferMatic(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	userId, ok := c.Get("userId").(int)
	if !ok {
		return makeErrorResponse(c, http.StatusUnauthorized, errors.New("invalid user id"))
	}

	req := new(TransferMaticRequest)

	if err := c.Bind(req); err != nil {
		return makeErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := h.service.TransferMatic(ctx, userId, req.ReceiverId, req.Amount); err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

// @Summary Transfer Ruble
// @Security ApiKeyAuth
// @Tags finance
// @Description transfer ruble from current user to another
// @Accept json
// @Produce json
// @Param transferRuble body TransferRubleRequest true "transferRuble"
// @Success 200
// @Failure 401 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/transfer/ruble [post]
func (h *Handler) TransferRuble(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	userId, ok := c.Get("userId").(int)
	if !ok {
		return makeErrorResponse(c, http.StatusUnauthorized, errors.New("invalid user id"))
	}

	req := new(TransferRubleRequest)

	if err := c.Bind(req); err != nil {
		return makeErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := h.service.TransferRuble(ctx, userId, req.ReceiverId, req.Amount); err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

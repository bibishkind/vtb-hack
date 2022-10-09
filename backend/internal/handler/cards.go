package handler

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateCardRequest struct {
	Card *domain.Card `json:"card"`
}

type CreateCardResponse struct {
	CardId int `json:"cardId"`
}

// @Summary Create Card
// @Security ApiKeyAuth
// @Tags cards
// @Description create card
// @Accept json
// @Produce json
// @Param card body CreateCardRequest true "card"
// @Success 200 {object} CreateCardResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/cards [post]
func (h *Handler) CreateCard(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	userId, ok := c.Get("userId").(int)
	if !ok {
		return makeErrorResponse(c, http.StatusUnauthorized, errors.New("invalid user id"))
	}

	req := new(CreateCardRequest)

	if err := c.Bind(req); err != nil {
		return makeErrorResponse(c, http.StatusBadRequest, err)
	}

	cardId, err := h.service.CreateCard(ctx, userId, req.Card)
	if err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"cardId": cardId,
	})
}

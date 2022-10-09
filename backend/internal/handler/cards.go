package handler

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CreateCardRequest struct {
	Card *domain.Card `json:"card"`
}

type CreateCardResponse struct {
	CardId int `json:"cardId"`
}

type GetAllCardsResponse struct {
	Cards []*domain.Card `json:"cards"`
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

// @Summary Get All Cards
// @Tags cards
// @Description get all cards
// @Accept json
// @Produce json
// @Success 200 {object} GetAllCardsResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/cards [get]
func (h *Handler) GetAllCards(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	cards, err := h.service.GetAllCards(ctx)
	if err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"cards": cards,
	})
}

// @Summary Delete Card
// @Security ApiKeyAuth
// @Tags cards
// @Description delete card
// @Param card_id   path int true "Card Id"
// @Accept json
// @Produce json
// @Success 204
// @Failure 401 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/cards/{card_id} [delete]
func (h *Handler) DeleteCard(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	userId, ok := c.Get("userId").(int)
	if !ok {
		return makeErrorResponse(c, http.StatusUnauthorized, errors.New("invalid user id"))
	}

	cardIdString := c.Param("card_id")
	cardId, err := strconv.Atoi(cardIdString)
	if err != nil {
		return makeErrorResponse(c, http.StatusBadRequest, err)
	}

	if err = h.service.DeleteCard(ctx, userId, cardId); err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

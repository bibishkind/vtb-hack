package handler

import (
	"coffee-layered-architecture/internal/domain"
)

type CreateCardRequest struct {
	Card *domain.Card `json:"card"`
}

type CreateCardResponse struct {
	CardId int `json:"cardId"`
}

// @Summary Create Card
// @Security ApiKeyAuth
// @Tags finance
// @Description create card
// @Accept json
// @Produce json
// @Param card body CreateCardRequest true "card"
// @Success 200 {object} CreateCardResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/cards [post]
//func (h *Handler) CreateCard(c echo.Context) error {
//	ctx := context.Background()
//	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)
//
//	userId, ok := c.Get("userId").(int)
//	if !ok {
//		return makeErrorResponse(c, http.StatusUnauthorized, errors.New("invalid user id"))
//	}
//
//	cardId, err := h.service.CreateCard(ctx, userId)
//	if err != nil {
//		return makeErrorResponse(c, http.StatusInternalServerError, err)
//	}
//
//	return c.JSON(http.StatusOK, map[string]interface{}{
//		"cardId": cardId,
//	})
//}

package handler

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UpdateScoreRequest struct {
	Score int `json:"score"`
}

// @Summary Update Score
// @Security ApiKeyAuth
// @Tags scores
// @Description update score
// @Accept json
// @Produce json
// @Param updateScore body UpdateScoreRequest true "updateScore"
// @Success 200
// @Failure 401 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/score [put]
func (h *Handler) UpdateScore(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	userId, ok := c.Get("userId").(int)
	if !ok {
		return makeErrorResponse(c, http.StatusUnauthorized, errors.New("invalid user id"))
	}

	req := new(UpdateScoreRequest)

	if err := c.Bind(req); err != nil {
		return makeErrorResponse(c, http.StatusBadRequest, err)
	}

	if err := h.service.UpdateScore(ctx, userId, req.Score); err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

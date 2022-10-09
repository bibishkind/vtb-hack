package handler

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetProfileResponse struct {
	Profile *domain.Profile `json:"profile"`
}

// @Summary Get Profile
// @Security ApiKeyAuth
// @Tags profile
// @Description get profile
// @Accept json
// @Produce json
// @Success 200 {object} GetProfileResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/profile [get]
func (h *Handler) GetProfile(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	userId, ok := c.Get("userId").(int)
	if !ok {
		return makeErrorResponse(c, http.StatusUnauthorized, errors.New("invalid user id"))
	}

	profile, err := h.service.GetProfile(ctx, userId)
	if err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"profile": profile,
	})
}

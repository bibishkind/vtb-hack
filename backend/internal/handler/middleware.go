package handler

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (h *Handler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

		authHeader := c.Request().Header.Get("Authorization")

		if len(authHeader) == 0 {
			return makeErrorResponse(c, http.StatusUnauthorized, errors.New("empty auth header"))
		}

		authHeaderSlice := strings.Split(authHeader, " ")

		if len(authHeaderSlice) != 2 {
			return makeErrorResponse(c, http.StatusUnauthorized, errors.New("invalid auth header"))
		}

		accessToken := authHeaderSlice[1]
		userId, err := h.service.ParseAccessToken(accessToken)
		if err != nil {
			return makeErrorResponse(c, http.StatusUnauthorized, err)
		}

		c.Set("userId", userId)
		return next(c)
	}
}

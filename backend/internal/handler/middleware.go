package handler

import (
	"context"
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
			return makeResponse(c, http.StatusUnauthorized, "empty auth header")
		}

		authHeaderSlice := strings.Split(authHeader, " ")

		if len(authHeaderSlice) != 2 {
			return makeResponse(c, http.StatusUnauthorized, "invalid authorization header")
		}

		accessToken := authHeaderSlice[1]
		userId, err := h.service.IdentifyUser(ctx, accessToken)
		if err != nil {
			return makeResponse(c, http.StatusUnauthorized, "failed to identify user")
		}

		c.Set("userId", userId)

		return nil
	}
}

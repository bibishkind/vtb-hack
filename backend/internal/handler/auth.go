package handler

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	AccessToken string `json:"accessToken"`
}

// @Summary Authorization of a user
// @Tags auth
// @Description authorizes user
// @Accept json
// @Produce json
// @Param user body SignUpRequest true "user"
// @Success 201
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return makeErrorResponse(c, http.StatusBadRequest, err)
	}

	if user.Username == "" || user.Password == "" {
		return makeErrorResponse(c, http.StatusBadRequest, errors.New("username and password required"))
	}

	if err := h.service.SignUp(ctx, user); err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusCreated)
}

// @Summary Authentication of a user
// @Tags auth
// @Description authenticates user
// @Accept json
// @Produce json
// @Param user body SignInRequest true "user"
// @Success 200 {object} SignInResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return makeErrorResponse(c, http.StatusBadRequest, err)
	}

	if user.Username == "" || user.Password == "" {
		return makeErrorResponse(c, http.StatusBadRequest, errors.New("username and password required"))
	}

	accessToken, err := h.service.SignIn(ctx, user)
	if err != nil {
		return makeErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": accessToken,
	})
}

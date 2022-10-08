package handler

import (
	"coffee-layered-architecture/internal/domain"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SwaggerSignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SwaggerSignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SwaggerSignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// @Summary Authorization of a user
// @Tags auth
// @Description authorizes user
// @Accept json
// @Produce json
// @Param user body SwaggerSignUpRequest true "user"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return h.makeResponse(c, http.StatusBadRequest, err.Error())
	}

	if user.Username == "" || user.Password == "" {
		return h.makeResponse(c, http.StatusBadRequest, "username and password are required")
	}

	if err := h.service.SignUp(ctx, user); err != nil {
		return h.makeResponse(c, http.StatusInternalServerError, "can't sign up the user")
	}

	return h.makeResponse(c, http.StatusCreated, "ok")
}

// @Summary Authentication of a user
// @Tags auth
// @Description authenticates user
// @Accept json
// @Produce json
// @Param user body SwaggerSignInRequest true "user"
// @Success 200 {object} SwaggerSignInResponse
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, h.requestTimeout)

	user := new(domain.User)
	if err := c.Bind(&user); err != nil {
		return h.makeResponse(c, http.StatusBadRequest, "can't bind the user")
	}

	accessToken, refreshToken, err := h.service.SignIn(ctx, user)
	if err != nil {
		return h.makeResponse(c, http.StatusInternalServerError, "can't sign in the user")
	}

	return c.JSON(http.StatusOK, struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

package handler

import (
	"coffee-layered-architecture/internal/config"
	service2 "coffee-layered-architecture/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"time"
)

type Handler struct {
	service        service2.Client
	requestTimeout time.Duration
}

func NewHandler(cfg *config.Config, service service2.Client) *Handler {
	return &Handler{
		service:        service,
		requestTimeout: cfg.RequestTimeout,
	}
}

func (h *Handler) Init() *echo.Echo {
	router := echo.New()

	api := router.Group("/api")
	{
		api.POST("/auth/sign-up", h.SignUp)
		api.POST("/auth/sign-in", h.SignIn)
	}

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	return router
}

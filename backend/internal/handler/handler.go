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
		requestTimeout: cfg.Handler.RequestTimeout,
	}
}

func (h *Handler) Init() *echo.Echo {
	router := echo.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	//api := router.Group("/api", h.AuthMiddleware)
	//{
	//
	//}

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	return router
}

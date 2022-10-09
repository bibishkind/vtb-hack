package handler

import (
	"coffee-layered-architecture/internal/config"
	service2 "coffee-layered-architecture/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"time"
)

type Handler struct {
	service        service2.Service
	requestTimeout time.Duration
}

func NewHandler(cfg *config.Config, service service2.Service) *Handler {
	return &Handler{
		service:        service,
		requestTimeout: cfg.Handler.RequestTimeout,
	}
}

func (h *Handler) Init() *echo.Echo {
	router := echo.New()

	router.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api")
	{
		api.GET("/balance", h.GetBalance, h.AuthMiddleware)
		transfer := api.Group("/transfer")
		{
			transfer.POST("/matic", h.TransferMatic, h.AuthMiddleware)
			transfer.POST("/ruble", h.TransferRuble, h.AuthMiddleware)
		}

		api.GET("/cards", h.GetAllCards)
		api.POST("/cards", h.CreateCard, h.AuthMiddleware)
		api.DELETE("/cards/:card_id", h.DeleteCard, h.AuthMiddleware)

		api.POST("/tasks", h.CreateTask, h.AuthMiddleware)
		api.GET("/tasks", h.GetAllTasks)

		api.GET("/profile", h.GetProfile, h.AuthMiddleware)

		api.PUT("/score", h.UpdateScore, h.AuthMiddleware)
	}

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	return router
}

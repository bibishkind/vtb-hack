package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateWallet(c echo.Context) error {
	uid := c.Param("userId")
	fmt.Println(uid)
	return nil
}

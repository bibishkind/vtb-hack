package handler

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func makeErrorResponse(c echo.Context, code int, err error) error {
	return c.JSON(code, &ErrorResponse{
		Code: code,
		Msg:  err.Error(),
	})
}

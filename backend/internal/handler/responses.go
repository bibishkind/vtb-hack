package handler

import "github.com/labstack/echo/v4"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func makeResponse(c echo.Context, code int, msg string) error {
	return c.JSON(code, &Response{
		Code: code,
		Msg:  msg,
	})
}

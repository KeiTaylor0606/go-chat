package main

import (
	"go-chat/handler"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/message", handler.PostMessageHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

package handler

import (
	"go-chat/db"
	"net/http"

	"github.com/labstack/echo"
)

func PostMessageHandler(c echo.Context) error {
	var params struct {
		teacherId uint
		studentId uint
		message   string
	}
	if err := c.Bind(params); err != nil {
		return err
	}
	err := db.Create(uint(params.teacherId), uint(params.studentId), params.message)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, params.message+"\n")
}

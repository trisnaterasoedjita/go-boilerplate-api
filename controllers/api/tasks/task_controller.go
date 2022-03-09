package tasks

import (
	"github.com/labstack/echo"
	"net/http"
)

func Index(c echo.Context) error {
	
	return c.JSON(http.StatusOK, map[string]string{
			"message": "all tasks data",
		})
}
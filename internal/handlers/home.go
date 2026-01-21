package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

// HomeHandler обрабатывает главную страницу
func HomeHandler(c *echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Echo Starter!")
}

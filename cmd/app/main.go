package main

import (
	"echo_starter/internal/handlers"
	"echo_starter/internal/middleware"

	"github.com/labstack/echo/v5"
)

// Основная функция приложения
func main() {
	e := echo.New()

	// Применяем кастомные middleware
	for _, m := range middleware.CustomMiddleware() {
		e.Use(m)
	}

	e.Static("/static", "static")

	// Регистрируем маршруты
	handlers.RegisterRoutes(e)

	// 2. ВАЖНО: Строгая привязка к IPv4
	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

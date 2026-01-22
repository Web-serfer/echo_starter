package middleware

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// CustomMiddleware возвращает набор кастомных middleware
func CustomMiddleware() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		// Логирование запросов
		middleware.RequestLogger(),

		// Восстановление после паники
		middleware.Recover(),

		// CORS с настройками
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"}, // Разрешить все домены
			AllowMethods: []string{
				"GET", "POST", "PUT", "DELETE", "OPTIONS",
			},
			AllowHeaders: []string{
				"Accept", "Authorization", "Content-Type", "X-CSRF-Token",
			},
		}),

		// Заголовки безопасности
		middleware.Secure(),
	}
}

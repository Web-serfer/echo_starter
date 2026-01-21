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
		
		// CORS
		middleware.CORS(),
		
		// Заголовки безопасности
		middleware.Secure(),
	}
}
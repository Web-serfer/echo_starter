package handlers

import (
	"echo_starter/view/pages"

	"github.com/labstack/echo/v5"
	"echo_starter/internal/render"
)

// HomeHandler обрабатывает главную страницу
func HomeHandler(c *echo.Context) error {
	component := pages.Index("Гость")
	return render.RenderTempl(c, 200, component)
}

// AboutHandler обрабатывает страницу "О нас"
func AboutHandler(c *echo.Context) error {
	component := pages.About("About Page")
	return render.RenderTempl(c, 200, component)
}

// ContactHandler обрабатывает страницу contacts
func ContactHandler(c *echo.Context) error {
	component := pages.Contact()
	return render.RenderTempl(c, 200, component)
}

// RegisterRoutes регистрирует все маршруты приложения
func RegisterRoutes(e *echo.Echo) {
	// Главная
	e.GET("/", HomeHandler)

	// О нас
	e.GET("/about", AboutHandler)

	// Контакты
	e.GET("/contact", ContactHandler)
}

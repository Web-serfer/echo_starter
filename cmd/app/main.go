package main

import (
	"echo_starter/view/pages"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// Основная функция приложения
func main() {
	e := echo.New()

	// 1. ВАЖНО: Recover восстанавливает сервер после паники
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLogger())

	e.Static("/static", "static")

	// Главная
	e.GET("/", func(c *echo.Context) error {
		// Используем хелпер (см. ниже)
		return render(c, pages.Index("Гость"))
	})

	// О нас
	e.GET("/about", func(c *echo.Context) error {
		return render(c, pages.About("About Page"))
	})

	// Контакты
	e.GET("/contact", func(c *echo.Context) error {
		return render(c, pages.Contact())
	})

	// 2. ВАЖНО: Строгая привязка к IPv4
	if err := e.Start("127.0.0.1:8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

// Хелпер для правильного рендеринга Templ в Echo
func render(c *echo.Context, component templ.Component) error {
	// 3. ВАЖНО: Явно задаем HTML, чтобы Air понимал, куда вставлять скрипт
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Рендерим сразу в Writer (быстрее и надежнее)
	return component.Render(c.Request().Context(), c.Response())
}

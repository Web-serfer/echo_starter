package main

import (
	"context"
	"net/http"
	"strings"

	"echo_starter/view/pages"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// Основная функция приложения
func main() {
	// Создаем новый экземпляр веб-фреймворка Echo
	e := echo.New()

	// Добавляем middleware для логирования входящих запросов
	e.Use(middleware.RequestLogger())

	// Добавляем статические файлы
	e.Static("/static", "static")

	// Определяем обработчик для главной страницы "/"
	e.GET("/", func(c *echo.Context) error {
		// Создаем компонент index с параметрами
		component := pages.Index("Гость")

		// Создаем строковый билдер для рендеринга компонента
		var sb strings.Builder
		err := component.Render(context.Background(), &sb)
		if err != nil {
			return err
		}

		// Возвращаем результат как HTML
		return c.HTML(http.StatusOK, sb.String())
	})

	// Определяем обработчик для страницы "О нас"
	e.GET("/about", func(c *echo.Context) error {
		// Создаем компонент about с параметрами
		component := pages.About("About Page")

		// Создаем строковый билдер для рендеринга компонента
		var sb strings.Builder
		err := component.Render(context.Background(), &sb)
		if err != nil {
			return err
		}

		// Возвращаем результат как HTML
		return c.HTML(http.StatusOK, sb.String())
	})

	// Определяем обработчик для страницы "Контакты"
	e.GET("/contact", func(c *echo.Context) error {
		// Создаем компонент contact с параметрами
		component := pages.Contact()

		// Создаем строковый билдер для рендеринга компонента
		var sb strings.Builder
		err := component.Render(context.Background(), &sb)
		if err != nil {
			return err
		}

		// Возвращаем результат как HTML
		return c.HTML(http.StatusOK, sb.String())
	})

	// Запускаем сервер на порту 8080
	// Если произойдет ошибка при запуске, выводим ее в лог
	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

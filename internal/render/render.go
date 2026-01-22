package render

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
)

// RenderTempl рендерит templ компонент
func RenderTempl(c *echo.Context, statusCode int, component templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	c.Response().WriteHeader(statusCode)
	return component.Render(c.Request().Context(), c.Response())
}
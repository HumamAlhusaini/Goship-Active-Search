package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, status int, com templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := com.Render(c.Request().Context(), buf); err != nil {
		return err
	}

	return c.HTML(status, buf.String())
}

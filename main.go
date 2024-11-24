package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/haatos/active-search-example/internal/handler"
	"github.com/haatos/active-search-example/internal/views/examples"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return handler.Render(c, http.StatusOK, examples.ActiveSearchExampleTable())
	})
	e.GET("/active-search", handler.GetActiveSearchExample)
	e.Start(":8080")
}

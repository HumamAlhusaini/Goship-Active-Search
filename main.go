package main

import (
	"github.com/labstack/echo/v4"

	"github.com/HumamAlhusaini/Goship-Active-Search/internal/handler"
)

func main() {
	e := echo.New()

	e.GET("/", handler.GetIndexPage)

	e.GET("/active-search", handler.GetActiveSearchExample)
	e.Start(":8080")
}

package handler

import (
	"net/http"

	"github.com/HumamAlhusaini/Goship-Active-Search/internal/views/pages"
	"github.com/labstack/echo/v4"
)

func GetIndexPage(c echo.Context) error {
	return render(c, http.StatusOK, pages.IndexPage())
}

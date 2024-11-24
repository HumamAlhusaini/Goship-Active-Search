package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/haatos/active-search-example/internal/views/examples"
	"github.com/labstack/echo/v4"
)

func GetActiveSearchExample(c echo.Context) error {
	time.Sleep(500 * time.Millisecond)

	search := strings.ToLower(strings.TrimSpace(c.FormValue("search")))

	out := make([]templ.Component, 0, len(examples.ActiveSearchTableData))

	for _, rowData := range examples.ActiveSearchTableData {
		if search == "" || (strings.Contains(strings.ToLower(rowData.FirstName), search) ||
			strings.Contains(strings.ToLower(rowData.LastName), search) ||
			strings.Contains(strings.ToLower(rowData.Email), search)) {
			out = append(out, examples.ActiveSearchTableRow(
				rowData.FirstName, rowData.LastName, rowData.Email))
		}
	}

	return Render(c, http.StatusOK, examples.ActiveSearchTableRows(out))
}

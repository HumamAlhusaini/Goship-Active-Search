package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/HumamAlhusaini/Goship-Active-Search/internal/model"
	"github.com/HumamAlhusaini/Goship-Active-Search/internal/views/components"
	"github.com/HumamAlhusaini/Goship-Active-Search/internal/views/examples"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// ActiveSearchExampleTable
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

	return render(c, http.StatusOK, examples.ActiveSearchTableRows(out))
}

func GetActiveSearchExampleTable(c echo.Context) error {
	out := make([]templ.Component, 0, len(examples.ActiveSearchTableData))
	for _, rowData := range examples.ActiveSearchTableData {
		out = append(out, examples.ActiveSearchTableRow(
			rowData.FirstName, rowData.LastName, rowData.Email))
	}

	com := examples.ActiveSearchExample(
		"active-search-table",
		[]templ.Component{
			components.PlainText("First Name"),
			components.PlainText("Last Name"),
			components.PlainText("Email"),
		},
		out)

	return render(c, http.StatusOK, com)
}

// ActiveSearchExampleTable

// InfiniteScrollTableExample
func GetInfiniteScrollExample(c echo.Context) error {
	// generate dummy row data
	data := make([][]string, 0, 10)
	for i := 0; i < 10; i++ {
		data = append(data, []string{"John Doe", fmt.Sprintf("john.doe%d@email.com", i)})
	}

	hasMore := true

	rows := make([]templ.Component, 0, len(data))
	for i := range data {
		rows = append(rows, components.InfiniteScrollRow(data[i][0], data[i][1], 1, i+1 == 10 && hasMore))
	}

	return render(c, http.StatusOK, components.InfiniteScrollTable(rows))
}

// LazyLoadExample
func GetLazyLoadExample(c echo.Context) error {
	time.Sleep(2 * time.Second)

	return render(c, http.StatusOK, examples.LazyLoadResult())
}

// LazyLoadExample

// PricingExample
func GetPricingExample(c echo.Context) error {
	yearly := c.QueryParam("period") == "on"

	return render(c, http.StatusOK, components.PriceGrid(examples.PriceDataExample(yearly)))
}

// PricingExample

// CascadingSelect
var modelOptions = map[string][]model.SelectOption{
	"audi": {
		{Label: "A1", Value: "a1"},
		{Label: "A4", Value: "a4"},
		{Label: "A6", Value: "a6"},
	},
	"bmw": {
		{Label: "316i", Value: "316i"},
		{Label: "320d", Value: "320d"},
	},
	"toyota": {
		{Label: "Corolla", Value: "corolla"},
		{Label: "Yaris", Value: "yaris"},
		{Label: "RAV4", Value: "rav4"},
	},
}

// CascadingSelect

// BasicPaginationExample
const (
	ItemsPerPage = 10
	PagesPerSide = 3
)

func getPaginationLowAndHigh(page, maxPages, pagePerSide int) (int, int) {
	low := max(0, page-pagePerSide-1)
	high := min(maxPages-1, page+pagePerSide-1)

	// if there are less than 'pagePerSide' pages to the left
	// of the current page, we add more pages to the high end of the
	// pagination element by adding to 'high'
	add := pagePerSide - page
	if add >= 0 {
		high += add + 1
	}

	// if there are less than 'pagePerSide' pages to the right
	// of the current page, we add more pages to the low end of the
	// pagination element, by subtracting from 'low'
	distanceFromMaxPages := maxPages - page
	if distanceFromMaxPages < pagePerSide {
		low -= (pagePerSide - distanceFromMaxPages)
	}
	return low, high
}

// BasicPaginationExample

// BasicCombobox
func PostCombobox(c echo.Context) error {
	name := c.Param("name")
	value := c.Param("value")

	return render(c, http.StatusOK, components.ComboBadge(name, value))
}

// BasicCombobox

// ModalConfirmDelete
// endpoint to log deleted 'element' and return an empty string
// to replace the element, i.e. remove it from the DOM
func DeleteModalExample(c echo.Context) error {
	value := c.QueryParam("value")
	log.Println("deleting modal", value)
	return c.String(http.StatusOK, "")
}

// ModalConfirmDelete

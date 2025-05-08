package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/smnschmnck/roundest-go-htmx/pages/results"
)

func RegisterPages(e *echo.Echo) {
	e.GET("/results", results.Page)
	e.POST("/vote", voteAction)
	e.GET("/", page)
}

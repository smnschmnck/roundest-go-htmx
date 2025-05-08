package results

import (
	"net/http"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/labstack/echo/v4"
	"github.com/smnschmnck/roundest-go-htmx/pages/layout"
	"github.com/smnschmnck/roundest-go-htmx/utils"
)

func Page(c echo.Context) error {
	page := layout.Layout(
		h.H1(
			g.Text("Results (TODO)"),
		),
	)

	return c.HTML(http.StatusOK, utils.RenderPage(page))
}

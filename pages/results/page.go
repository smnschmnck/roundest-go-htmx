package results

import (
	"fmt"
	"net/http"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/labstack/echo/v4"
	"github.com/smnschmnck/roundest-go-htmx/pages/layout"
	components "github.com/smnschmnck/roundest-go-htmx/pages/results/_components"
	resultUtils "github.com/smnschmnck/roundest-go-htmx/pages/results/utils"
	"github.com/smnschmnck/roundest-go-htmx/utils"
)

func th(label string) g.Node {
	return h.Th(
		h.Class("px-6 py-3 text-left text-xs font-medium text-white uppercase tracking-wider"),
		g.Text(label),
	)
}

func tHead() g.Node {
	return h.THead(
		h.Class("bg-gray-800"),
		th("Rank"),
		th("Pokemon"),
		th("Win Rate"),
		th("Loss Rate"),
		th("Total Votes"),
	)
}

func Page(c echo.Context) error {
	results, err := resultUtils.GetResults()
	if err != nil {
		fmt.Println(err.Error())
	}

	page := layout.Layout(
		h.Main(
			h.Class("container mx-auto px-4 py-8"),
			h.Div(
				h.Class("rounded-lg shadow overflow-hidden"),
				h.Table(
					h.Class("min-w-full divide-y divide-gray-700"),
					tHead(),
					h.TBody(
						h.Class("divide-y divide-gray-700"),
						g.Map(results, func(r resultUtils.ReadableResult) g.Node {
							return components.ResultRow(r)
						}),
					),
				),
			),
		),
	)

	return c.HTML(http.StatusOK, utils.RenderPage(page))
}

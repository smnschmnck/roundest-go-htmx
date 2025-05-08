package results

import (
	"context"
	"net/http"
	"strconv"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/labstack/echo/v4"
	"github.com/smnschmnck/roundest-go-htmx/db"
	"github.com/smnschmnck/roundest-go-htmx/db/queries"
	"github.com/smnschmnck/roundest-go-htmx/pages/layout"
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

func resultRow(idx int, results queries.GetResultsRow) g.Node {
	return h.Tr(
		h.Class("hover:bg-gray-800/50"),
		h.Td(
			h.Class("px-6 py-4 whitespace-nowrap text-sm text-white"),
			g.Text(strconv.Itoa(idx)),
		),
		h.Td(
			h.Class("px-6 py-4 whitespace-nowrap"),
			g.Text(results.Name),
		),
		h.Td(
			h.Class("px-6 py-4 whitespace-nowrap"),
			g.Text(strconv.Itoa(int(results.VotesFor))),
		),
		h.Td(
			h.Class("px-6 py-4 whitespace-nowrap"),
			g.Text(strconv.Itoa(int(results.VotesAgainst))),
		),
		h.Td(
			h.Class("px-6 py-4 whitespace-nowrap text-sm text-white"),
			g.Text(strconv.Itoa(int(results.VotesAgainst+results.VotesFor))),
		),
	)
}

func Page(c echo.Context) error {
	results, _ := db.DB.GetResults(context.Background())

	page := layout.Layout(
		h.Div(
			h.Class("container mx-auto px-4 py-8"),
			h.Div(
				h.Class("rounded-lg shadow overflow-hidden"),
				h.Table(
					h.Class("min-w-full divide-y divide-gray-700"),
					tHead(),
					h.TBody(
						h.Class("divide-y divide-gray-700"),
						g.Map(results, func(r queries.GetResultsRow) g.Node {
							return resultRow(1, r)
						}),
					),
				),
			),
		),
	)

	return c.HTML(http.StatusOK, utils.RenderPage(page))
}

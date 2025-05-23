package layout

import (
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func Layout(children ...g.Node) g.Node {
	return h.Div(
		h.Class("flex flex-col min-h-full h-full w-full"),
		h.Nav(
			hx.Boost("true"),
			h.Class("flex justify-between items-center px-8 py-4 border-t-2 border-t-blue-500"),
			h.Div(
				h.Class("flex items-center gap-2"),
				h.A(
					h.Href("/"),
					h.Class("text-3xl font-bold"),
					h.Span(
						h.Class("text-white"),
						g.Text("round"),
					),
					h.Span(
						h.Class("text-blue-500"),
						g.Text("est"),
					),
				),
				h.Span(
					h.Class("text-gray-400 text-2xl font-extralight"),
					g.Text("(Go/HTMX)"),
				),
			),
			h.A(
				h.Class("text-lg hover:underline"),
				h.Href("/results"),
				g.Text("Results"),
			),
		),
		g.Group(children),
		h.Footer(
			h.Class("flex w-full justify-center items-center"),
			h.A(
				h.Target("_blank"),
				h.Href("https://github.com/smnschmnck/roundest-go-htmx"),
				h.Class("font-light text-center py-3 text-gray-500"),
				g.Text("GitHub"),
			),
		),
	)
}

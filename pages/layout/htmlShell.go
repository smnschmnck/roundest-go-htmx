package layout

import (
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

func HtmlShell(children ...g.Node) g.Node {
	return c.HTML5(
		c.HTML5Props{
			Title:    "Roundest",
			Language: "en",
			Head: []g.Node{
				h.Link(h.Rel("stylesheet"), h.Href("/static/styles/app.css"), h.Type("text/css")),
				h.Link(h.Rel("icon"), h.Href("https://fav.farm/🍪")),
				h.Script(
					h.Src("https://unpkg.com/htmx.org@2.0.4/dist/htmx.min.js"),
					h.CrossOrigin("anonymous"),
				),
				h.Script(
					h.Src("https://unpkg.com/htmx-ext-preload@2.1.0/preload.js"),
					h.CrossOrigin("anonymous"),
				),
			},
			HTMLAttrs: []g.Node{h.Class("h-full w-full")},
			Body: []g.Node{
				hx.Ext("preload"),
				g.Group(children),
			},
		})
}

package layout

import (
	g "maragu.dev/gomponents"
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
				h.Script(h.Src("https://unpkg.com/htmx.org")),
			},
			HTMLAttrs: []g.Node{h.Class("h-full w-full")},
			Body:      children,
		})
}

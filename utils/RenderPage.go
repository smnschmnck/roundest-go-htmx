package utils

import (
	"bytes"

	g "maragu.dev/gomponents"
)

func RenderPage(page g.Node) string {
	renderedPage := new(bytes.Buffer)
	page.Render(renderedPage)
	return renderedPage.String()
}

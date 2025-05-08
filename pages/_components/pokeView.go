package components

import (
	"fmt"
	"strconv"

	"github.com/smnschmnck/roundest-go-htmx/db/queries"
	"github.com/smnschmnck/roundest-go-htmx/utils"
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func voteButton(winnerId int32, loserId int32) g.Node {
	winnerIdString := strconv.Itoa(int(winnerId))
	loserIdString := strconv.Itoa(int(loserId))

	return h.Form(
		hx.Post("/vote"),
		hx.Target("#battleground"),
		hx.Swap("outerHTML"),
		h.Input(h.Type("hidden"), h.Value(winnerIdString), h.Name("winnerId")),
		h.Input(h.Type("hidden"), h.Value(loserIdString), h.Name("loserId")),
		h.Button(
			h.Type("submit"),
			h.Class("hover:bg-gray-700 bg-blue-600 text-white px-4 py-2 rounded-md flex w-full justify-center items-centers"),
			g.Text("Vote"),
		),
	)
}

func PokeView(pokemon queries.Pokemon, loserId int32) g.Node {
	return h.Div(
		h.Class("flex flex-col gap-4"),
		h.Img(
			h.Class("w-48 h-48"),
			h.Style("image-rendering: pixelated;"),
			h.Src(utils.GetPokeImage(pokemon)),
		),
		h.Div(
			h.Class("flex flex-col justify-center items-center"),
			h.Span(
				h.Class("text-gray-500 text-lg"),
				g.Text(fmt.Sprintf("#%d", pokemon.ID)),
			),
			h.H2(
				h.Class("text-2xl font-bold capitalize"),
				g.Text(pokemon.Name),
			),
		),
		voteButton(pokemon.ID, loserId),
	)
}

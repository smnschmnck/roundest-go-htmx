package pages

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"

	"github.com/labstack/echo/v4"
	"github.com/smnschmnck/roundest-go-htmx/db"
	"github.com/smnschmnck/roundest-go-htmx/db/queries"
	components "github.com/smnschmnck/roundest-go-htmx/pages/_components"
	"github.com/smnschmnck/roundest-go-htmx/pages/layout"
	"github.com/smnschmnck/roundest-go-htmx/utils"
)

func getRandomPokemon() ([2]queries.Pokemon, error) {
	pokemon, err := db.DB.GetTwoRandomPokemon(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return [2]queries.Pokemon(pokemon), err
	}

	p1 := pokemon[0]
	p2 := pokemon[1]

	return [2]queries.Pokemon{p1, p2}, nil
}

func battleground() g.Node {
	pokemon, err := getRandomPokemon()
	if err != nil {
		return h.H1(
			g.Text("ERROR"),
		)
	}

	return h.Main(
		h.ID("battleground"),
		h.Class("w-full h-full flex items-center justify-center"),
		h.Div(
			h.Class("w-full h-full grow flex flex-col items-center justify-center gap-8"),
			h.Div(
				h.Class("md:grid grid-cols-2 gap-8"),
				components.PokeView(pokemon[0], pokemon[1].ID),
				components.PokeView(pokemon[1], pokemon[0].ID),
			),
		),
	)
}

func voteAction(c echo.Context) error {
	winnerId, err := strconv.Atoi(c.FormValue("winnerId"))
	if err != nil {
		fmt.Println(err.Error())
		return c.HTML(http.StatusOK, utils.RenderPage(battleground()))
	}
	loserId, err := strconv.Atoi(c.FormValue("loserId"))
	if err != nil {
		fmt.Println(err.Error())
		return c.HTML(http.StatusOK, utils.RenderPage(battleground()))
	}

	err = db.DB.InsertVote(context.Background(), queries.InsertVoteParams{
		VotedForID:     int32(winnerId),
		VotedAgainstID: int32(loserId),
	})
	if err != nil {
		fmt.Println(err.Error())
		return c.HTML(http.StatusOK, utils.RenderPage(battleground()))
	}

	return c.HTML(http.StatusOK, utils.RenderPage(battleground()))
}

func page(c echo.Context) error {
	page := layout.Layout(
		battleground(),
	)

	return c.HTML(http.StatusOK, utils.RenderPage(page))
}

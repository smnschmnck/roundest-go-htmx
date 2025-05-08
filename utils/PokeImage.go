package utils

import (
	"fmt"

	"github.com/smnschmnck/roundest-go-htmx/db/queries"
)

func GetPokeImageById(id int32) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/%d.png", id)
}

func GetPokeImage(pokemon queries.Pokemon) string {
	return GetPokeImageById(pokemon.ID)
}

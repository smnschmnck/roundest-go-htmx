package seed

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/smnschmnck/roundest-go-htmx/db/queries"
)

type PokemonResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type APIResponse struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous string          `json:"previous"`
	Results  []PokemonResult `json:"results"`
}

func fetchAllPokemon() ([]PokemonResult, error) {
	fmt.Println("FETCHING POKEMON...")

	url := "https://pokeapi.co/api/v2/pokemon?limit=1025&offset=0"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	fmt.Println("FETCHED POKEMON")

	return apiResp.Results, nil
}

var isSeeded = false

func Seed(db *queries.Queries) error {
	if isSeeded {
		fmt.Printf("DB IS SEEDED ALREADY. SKIPPING")
		return nil
	}

	dbIsSeeded, err := db.CheckIsSeeded(context.Background())
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	if dbIsSeeded {
		fmt.Printf("DB IS SEEDED ALREADY. SKIPPING")
		isSeeded = true
		return nil
	}

	fmt.Println("SEEDING...")

	allPokemon, err := fetchAllPokemon()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}

	fmt.Println("INSERTING POKEMON INTO DB...")

	for _, p := range allPokemon {
		splittedUrl := strings.Split(p.URL, "/")
		id := splittedUrl[len(splittedUrl)-2]
		idInteger, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("%s", err.Error())
			return err
		}
		db.CreatePokemon(context.Background(), queries.CreatePokemonParams{ID: int32(idInteger), Name: p.Name})
	}

	fmt.Println("SEEDED DB")

	return nil
}

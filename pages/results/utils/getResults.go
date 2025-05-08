package utils

import (
	"context"
	"fmt"
	"strconv"

	"github.com/smnschmnck/roundest-go-htmx/db"
	"github.com/smnschmnck/roundest-go-htmx/utils"
)

type ReadableResult struct {
	Rank           string
	ID             string
	Name           string
	ImgUrl         string
	TotalVotes     string
	Wins           string
	Losses         string
	WinPercentage  string
	LossPercentage string
}

func getPercentage(votes int64, total int64) string {
	if votes <= 0 {
		return "0"
	}
	if total <= 0 {
		return "0"
	}

	percentage := (float64(votes) / float64(total)) * 100

	return fmt.Sprintf("%.2f", percentage)
}

func GetResults() ([]ReadableResult, error) {
	results, err := db.DB.GetResults(context.Background())
	if err != nil {
		return make([]ReadableResult, 0), err
	}

	readableResults := make([]ReadableResult, 0)

	for idx, result := range results {
		totalVotes := result.VotesAgainst + result.VotesFor
		totalVotesString := strconv.Itoa(int(totalVotes))
		winPercentage := getPercentage(result.VotesFor, int64(totalVotes))
		lossPercentage := getPercentage(result.VotesAgainst, int64(totalVotes))

		readableResults = append(readableResults, ReadableResult{
			Rank:           strconv.Itoa(idx),
			ID:             fmt.Sprintf("#%d", result.ID),
			Name:           result.Name,
			ImgUrl:         utils.GetPokeImageById(result.ID),
			TotalVotes:     totalVotesString,
			Wins:           strconv.Itoa(int(result.VotesFor)),
			Losses:         strconv.Itoa(int(result.VotesAgainst)),
			WinPercentage:  winPercentage,
			LossPercentage: lossPercentage,
		})
	}

	return readableResults, nil
}

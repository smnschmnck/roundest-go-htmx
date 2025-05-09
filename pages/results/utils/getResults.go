package utils

import (
	"context"
	"fmt"
	"sort"
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

	sort.Slice(results, func(i, j int) bool {
		a := results[i]
		b := results[j]

		aRank := a.VotesFor - a.VotesAgainst
		bRank := b.VotesFor - b.VotesAgainst

		if aRank == bRank {
			return a.VotesFor > b.VotesFor
		}

		return aRank > bRank
	})

	readableResults := make([]ReadableResult, 0, len(results))

	for _, result := range results {
		totalVotes := result.VotesAgainst + result.VotesFor
		winPercentage := getPercentage(result.VotesFor, int64(totalVotes))
		lossPercentage := getPercentage(result.VotesAgainst, int64(totalVotes))

		readableResults = append(readableResults, ReadableResult{
			ID:             fmt.Sprintf("#%d", result.ID),
			Name:           result.Name,
			ImgUrl:         utils.GetPokeImageById(result.ID),
			TotalVotes:     strconv.Itoa(int(totalVotes)),
			Wins:           strconv.Itoa(int(result.VotesFor)),
			Losses:         strconv.Itoa(int(result.VotesAgainst)),
			WinPercentage:  winPercentage,
			LossPercentage: lossPercentage,
		})
	}

	for idx := range readableResults {
		readableResults[idx].Rank = strconv.Itoa(idx + 1) // Rank starts from 1
	}

	return readableResults, nil
}

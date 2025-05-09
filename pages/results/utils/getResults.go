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

func getPercentage(votes int64, total int64) float64 {
	if votes <= 0 {
		return 0
	}
	if total <= 0 {
		return 0
	}

	return (float64(votes) / float64(total)) * 100
}

func getPercentageString(votes int64, total int64) string {
	percentage := getPercentage(votes, total)

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

		aRank := getPercentage(a.VotesFor, a.VotesFor+a.VotesAgainst)
		bRank := getPercentage(b.VotesFor, b.VotesFor+b.VotesAgainst)

		if aRank == bRank {
			return a.VotesFor > b.VotesFor
		}

		return aRank > bRank
	})

	readableResults := make([]ReadableResult, 0, len(results))

	for _, result := range results {
		totalVotes := result.VotesAgainst + result.VotesFor
		winPercentage := getPercentageString(result.VotesFor, int64(totalVotes))
		lossPercentage := getPercentageString(result.VotesAgainst, int64(totalVotes))

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

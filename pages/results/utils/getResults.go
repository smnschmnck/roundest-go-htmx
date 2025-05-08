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

	sort.Slice(readableResults, func(i, j int) bool {
		aWins, err := strconv.Atoi(readableResults[i].Wins)
		if err != nil {
			return false
		}
		bWins, err := strconv.Atoi(readableResults[j].Wins)
		if err != nil {
			return false
		}
		aLosses, err := strconv.Atoi(readableResults[i].Losses)
		if err != nil {
			return false
		}
		bLosses, err := strconv.Atoi(readableResults[j].Losses)
		if err != nil {
			return false
		}

		aRank := aWins - aLosses
		bRank := bWins - bLosses

		return aRank > bRank

	})

	for idx := range readableResults {
		readableResults[idx].Rank = strconv.Itoa(idx + 1) // Rank starts from 1
	}

	return readableResults, nil
}

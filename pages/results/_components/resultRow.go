package components

import (
	"github.com/smnschmnck/roundest-go-htmx/pages/results/utils"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

const RESULT_INDICATOR_WINS = "wins"
const RESULT_INDICATOR_LOSSES = "losses"

func getPercentage(indicatorType string, result utils.ReadableResult) string {
	if indicatorType == RESULT_INDICATOR_WINS {
		return result.WinPercentage
	}
	if indicatorType == RESULT_INDICATOR_LOSSES {
		return result.LossPercentage
	}

	return "0"
}

func getVotes(indicatorType string, result utils.ReadableResult) string {
	if indicatorType == RESULT_INDICATOR_WINS {
		return result.Wins
	}
	if indicatorType == RESULT_INDICATOR_LOSSES {
		return result.Losses
	}

	return "0"
}

func getProgressbarClass(indicatorType string) string {
	if indicatorType == RESULT_INDICATOR_LOSSES {
		return "w-full h-2 [&::-webkit-progress-bar]:rounded-lg [&::-webkit-progress-value]:rounded-lg [&::-webkit-progress-bar]:bg-gray-700 [&::-webkit-progress-value]:bg-red-600 [&::-moz-progress-bar]:bg-red-600"
	}

	return "w-full h-2 [&::-webkit-progress-bar]:rounded-lg [&::-webkit-progress-value]:rounded-lg [&::-webkit-progress-bar]:bg-gray-700 [&::-webkit-progress-value]:bg-green-600 [&::-moz-progress-bar]:bg-green-600"
}

func resultIndicator(indicatorType string, result utils.ReadableResult) g.Node {
	if indicatorType != RESULT_INDICATOR_LOSSES && indicatorType != RESULT_INDICATOR_WINS {
		return h.Span(g.Text("Invalid indicator type"))
	}

	percentage := getPercentage(indicatorType, result)
	votes := getVotes(indicatorType, result)

	return h.Div(
		h.Class("flex flex-col w-48 gap-1"),
		h.Div(
			h.Class("flex justify-between"),
			h.Span(
				h.Class("text-sm font-medium text-white"),
				g.Text(percentage+"%"),
			),
			h.Span(
				h.Class("text-sm text-white/75"),
				g.Text(votes+" "+indicatorType),
			),
		),
		h.Progress(
			h.Class(getProgressbarClass(indicatorType)),
			h.Value(percentage),
			h.Max("100"),
		),
	)
}

func ResultRow(result utils.ReadableResult) g.Node {
	return h.Tr(
		h.Class("hover:bg-gray-800/50"),
		h.Td(
			h.Class("px-6 py-4 whitespace-nowrap text-sm text-white"),
			g.Text(result.Rank),
		),
		h.Td(
			h.Class("px-6 py-4 whitespace-nowrap"),
			h.Div(
				h.Class("flex items-center"),
				h.Img(
					h.Class("w-10 h-10 rounded-lg bg-gray-800 p-1"),
					h.Src(result.ImgUrl),
				),
				h.Div(
					h.Class("ml-4"),
					h.Div(
						h.Class("text-sm font-medium text-white"),
						g.Text(result.Name),
					),
					h.Div(
						h.Class("text-sm text-white/75"),
						g.Text(result.ID),
					),
				),
			),
		),
		h.Td(
			h.Class("px-6 py-4 whitespace-nowrap"),
			resultIndicator(RESULT_INDICATOR_WINS, result),
		),
		h.Td(
			h.Class("px-6 py-4 whitespace-nowrap"),
			resultIndicator(RESULT_INDICATOR_LOSSES, result),
		),
		h.Td(
			h.Class("px-6 py-4 whitespace-nowrap text-sm text-white"),
			g.Text(result.TotalVotes),
		),
	)
}

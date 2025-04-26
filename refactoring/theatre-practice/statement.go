package theatre

import (
	"fmt"
	"math"

	"github.com/leekchan/accounting"
)

type StatementPrinter struct {
	plays map[string]Play
}

func (s StatementPrinter) Print(invoice Invoice, plays map[string]Play) (string, error) {
	s.plays = plays

	totalAmount := 0
	volumeCredits := 0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)

	ac := accounting.Accounting{Symbol: "$", Precision: 2}

	for _, perf := range invoice.Performances {
		volumeCredits += s.volumeCreditsFor(perf)

		amount, err := s.amount(s.playFor(perf), perf)
		if err != nil {
			return "", err
		}
		// print line for this order
		result += fmt.Sprintf(
			"  %s: %s (%d seats)\n",
			s.playFor(perf).Name,
			ac.FormatMoney(float64(amount)/100),
			perf.Audience,
		)
		totalAmount += amount
	}
	result += fmt.Sprintf("Amount owed is %s\n", ac.FormatMoney(float64(totalAmount)/100))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)
	return result, nil
}

func (s StatementPrinter) volumeCreditsFor(perf Performance) int {
	result := 0
	result += int(math.Max(float64(perf.Audience)-30, 0))
	if s.playFor(perf).Type == "comedy" {
		result += int(math.Floor(float64(perf.Audience) / 5))
	}
	return result
}

func (s StatementPrinter) playFor(perf Performance) Play {
	return s.plays[perf.PlayID]
}

func (StatementPrinter) amount(play Play, perf Performance) (int, error) {
	result := 0

	switch play.Type {
	case "tragedy":
		result = 40000
		if perf.Audience > 30 {
			result += 1000 * (perf.Audience - 30)
		}
	case "comedy":
		result = 30000
		if perf.Audience > 20 {
			result += 10000 + 500*(perf.Audience-20)
		}
		result += 300 * perf.Audience
	default:
		return 0, fmt.Errorf("unknown type: %s", play.Type)
	}

	return result, nil
}

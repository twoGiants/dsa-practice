package theatre

import (
	"fmt"
	"math"

	"github.com/leekchan/accounting"
)

type StatementPrinter struct{}

func (s StatementPrinter) Print(invoice Invoice, plays map[string]Play) (string, error) {
	totalAmount := 0
	volumeCredits := 0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)

	ac := accounting.Accounting{Symbol: "$", Precision: 2}

	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		amount, err := s.amount(play, perf)
		if err != nil {
			return "", err
		}

		// add volume credits
		volumeCredits += int(math.Max(float64(perf.Audience)-30, 0))
		// add extra credit for every ten comedy attendees
		if play.Type == "comedy" {
			volumeCredits += int(math.Floor(float64(perf.Audience) / 5))
		}

		// print line for this order
		result += fmt.Sprintf(
			"  %s: %s (%d seats)\n",
			play.Name,
			ac.FormatMoney(float64(amount)/100),
			perf.Audience,
		)
		totalAmount += amount
	}
	result += fmt.Sprintf("Amount owed is %s\n", ac.FormatMoney(float64(totalAmount)/100))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)
	return result, nil
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

package theatre

import (
	"fmt"
	"math"

	"github.com/leekchan/accounting"
)

type StatementPrinter struct {
	plays   map[string]Play
	invoice Invoice
}

type StatementData struct {
	Customer     string
	Performances []EnrichedPerformance
	TotalAmount  int
}

type EnrichedPerformance struct {
	Name     string
	Amount   int
	Audience int
}

func (s StatementPrinter) Print(invoice Invoice, plays map[string]Play) (string, error) {
	s.plays = plays
	s.invoice = invoice

	statementData := StatementData{}
	statementData.Customer = invoice.Customer

	enrichedPerformances := []EnrichedPerformance{}
	for _, perf := range invoice.Performances {
		ep, err := s.enrichPerformance(perf)
		if err != nil {
			return "", err
		}
		enrichedPerformances = append(enrichedPerformances, ep)
	}
	statementData.Performances = enrichedPerformances
	statementData.TotalAmount = totalAmount(enrichedPerformances)

	return renderPlayText(s, statementData)
}

func (s StatementPrinter) enrichPerformance(perf Performance) (EnrichedPerformance, error) {
	result := EnrichedPerformance{}
	amount, err := s.amount(perf)
	if err != nil {
		return EnrichedPerformance{}, err
	}
	result.Amount = amount
	result.Audience = perf.Audience
	result.Name = s.playFor(perf).Name
	return result, nil
}

func renderPlayText(s StatementPrinter, data StatementData) (string, error) {
	result := fmt.Sprintf("Statement for %s\n", data.Customer)

	for _, perf := range data.Performances {
		result += fmt.Sprintf(
			"  %s: %s (%d seats)\n",
			perf.Name,
			usd(perf.Amount),
			perf.Audience,
		)
	}

	result += fmt.Sprintf("Amount owed is %s\n", usd(data.TotalAmount))
	result += fmt.Sprintf("You earned %d credits\n", s.totalVolumeCredits())

	return result, nil
}

func totalAmount(ep []EnrichedPerformance) int {
	result := 0
	for _, perf := range ep {
		result += perf.Amount
	}
	return result
}

func (s StatementPrinter) totalVolumeCredits() int {
	result := 0
	for _, perf := range s.invoice.Performances {
		result += s.volumeCreditsFor(perf)
	}
	return result
}

func usd(amount int) string {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	return ac.FormatMoney(float64(amount) / 100)
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

func (s StatementPrinter) amount(perf Performance) (int, error) {
	result := 0

	switch s.playFor(perf).Type {
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
		return 0, fmt.Errorf("unknown type: %s", s.playFor(perf).Type)
	}

	return result, nil
}

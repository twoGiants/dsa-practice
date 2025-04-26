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
	Customer           string
	Performances       []EnrichedPerformance
	TotalAmount        int
	TotalVolumeCredits int
	plays              map[string]Play
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
	statementData.plays = plays
	statementData.Customer = invoice.Customer

	enrichedPerformances := []EnrichedPerformance{}
	for _, perf := range invoice.Performances {
		ep, err := statementData.enrichPerformance(perf)
		if err != nil {
			return "", err
		}
		enrichedPerformances = append(enrichedPerformances, ep)
	}
	statementData.Performances = enrichedPerformances
	statementData.TotalAmount = statementData.totalAmount()
	statementData.TotalVolumeCredits = s.totalVolumeCredits()

	return renderPlayText(statementData)
}

func (s StatementData) enrichPerformance(perf Performance) (EnrichedPerformance, error) {
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

func (s StatementData) amount(perf Performance) (int, error) {
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

func (s StatementData) playFor(perf Performance) Play {
	return s.plays[perf.PlayID]
}

func renderPlayText(data StatementData) (string, error) {
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
	result += fmt.Sprintf("You earned %d credits\n", data.TotalVolumeCredits)

	return result, nil
}

func (s StatementData) totalAmount() int {
	result := 0
	for _, perf := range s.Performances {
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

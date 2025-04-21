package theatre

import (
	"fmt"
	"math"

	"github.com/leekchan/accounting"
)

type Play struct {
	Name string
	Type string
}

type Invoice struct {
	Customer     string
	Performances []Performance
}

type Performance struct {
	PlayID   string
	Audience int
}

type EnrichedPerformance struct {
	PlayID   string
	Audience int
	play     Play
	amount   int
}

type StatementPrinter struct {
	invoice Invoice
	plays   map[string]Play
}

func (s StatementPrinter) Print(invoice Invoice, plays map[string]Play) (string, error) {
	s.invoice = invoice
	s.plays = plays
	statementData := StatementData{s.invoice.Customer, nil}
	enrichedPerformances, err := s.enrichPerformances(s.invoice.Performances)
	if err != nil {
		return "", err
	}
	statementData.Performances = enrichedPerformances

	return PlainTextStatement{statementData, plays}.Render(), nil
}

func (s StatementPrinter) enrichPerformances(performances []Performance) ([]EnrichedPerformance, error) {
	var result []EnrichedPerformance
	for _, performance := range performances {
		enrichedPerformance := EnrichedPerformance{}
		enrichedPerformance.PlayID = performance.PlayID
		enrichedPerformance.Audience = performance.Audience
		enrichedPerformance.play = s.playFor(performance)

		amount, err := s.amountFor(enrichedPerformance)
		if err != nil {
			return nil, err
		}
		enrichedPerformance.amount = amount

		result = append(result, enrichedPerformance)
	}

	return result, nil
}

func (s StatementPrinter) playFor(aPerformance Performance) Play {
	return s.plays[aPerformance.PlayID]
}

func (StatementPrinter) amountFor(aPerformance EnrichedPerformance) (int, error) {
	result := 0
	switch aPerformance.play.Type {
	case "tragedy":
		result = 40000
		if aPerformance.Audience > 30 {
			result += 1000 * (aPerformance.Audience - 30)
		}
	case "comedy":
		result = 30000
		if aPerformance.Audience > 20 {
			result += 10000 + 500*(aPerformance.Audience-20)
		}
		result += 300 * aPerformance.Audience
	default:
		return 0, fmt.Errorf("unknown type: %s", aPerformance.play.Type)
	}
	return result, nil
}

type PlainTextStatement struct {
	data  StatementData
	plays map[string]Play
}

type StatementData struct {
	Customer     string
	Performances []EnrichedPerformance
}

func (s PlainTextStatement) Render() string {
	result := fmt.Sprintf("Statement for %s\n", s.data.Customer)

	for _, perf := range s.data.Performances {
		result += fmt.Sprintf(
			"  %s: %s (%d seats)\n",
			perf.play.Name,
			s.usd(perf.amount),
			perf.Audience,
		)
	}

	result += fmt.Sprintf("Amount owed is %s\n", s.usd(s.totalAmount()))
	result += fmt.Sprintf("You earned %d credits\n", s.totalVolumeCredits())
	return result
}

func (s PlainTextStatement) totalAmount() int {
	result := 0
	for _, perf := range s.data.Performances {
		result += perf.amount
	}
	return result
}

func (s PlainTextStatement) totalVolumeCredits() int {
	result := 0
	for _, perf := range s.data.Performances {
		result += s.volumeCreditsFor(perf)
	}
	return result
}

func (PlainTextStatement) usd(number int) string {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	return ac.FormatMoney(number / 100)
}

func (PlainTextStatement) volumeCreditsFor(aPerformance EnrichedPerformance) int {
	result := int(math.Max(float64(aPerformance.Audience)-30, 0))
	// add extra credit for every ten comedy attendees
	if aPerformance.play.Type == "comedy" {
		result += int(math.Floor(float64(aPerformance.Audience) / 5))
	}
	return result
}

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

type StatementData struct{}

type StatementPrinter struct{}

func (s StatementPrinter) Print(invoice Invoice, plays map[string]Play) (string, error) {
	statementData := StatementData{}
	return PlainTextStatement{plays, invoice, statementData}.Render()
}

type PlainTextStatement struct {
	plays   map[string]Play
	invoice Invoice
	data    StatementData
}

func (s PlainTextStatement) Render() (string, error) {
	result := fmt.Sprintf("Statement for %s\n", s.invoice.Customer)

	for _, perf := range s.invoice.Performances {
		amount, err := s.amountFor(perf)
		if err != nil {
			return "", err
		}
		result += fmt.Sprintf(
			"  %s: %s (%d seats)\n",
			s.playFor(perf).Name,
			s.usd(amount),
			perf.Audience,
		)
	}

	totalAmount, err := s.totalAmount()
	if err != nil {
		return "", err
	}
	result += fmt.Sprintf("Amount owed is %s\n", s.usd(totalAmount))
	result += fmt.Sprintf("You earned %d credits\n", s.totalVolumeCredits())
	return result, nil
}

func (s PlainTextStatement) totalAmount() (int, error) {
	result := 0
	for _, perf := range s.invoice.Performances {
		amount, err := s.amountFor(perf)
		if err != nil {
			return 0, err
		}
		result += amount
	}
	return result, nil
}

func (s PlainTextStatement) totalVolumeCredits() int {
	result := 0
	for _, perf := range s.invoice.Performances {
		result += s.volumeCreditsFor(perf)
	}
	return result
}

func (PlainTextStatement) usd(number int) string {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	return ac.FormatMoney(number / 100)
}

func (s PlainTextStatement) volumeCreditsFor(aPerformance Performance) int {
	result := int(math.Max(float64(aPerformance.Audience)-30, 0))
	// add extra credit for every ten comedy attendees
	if s.playFor(aPerformance).Type == "comedy" {
		result += int(math.Floor(float64(aPerformance.Audience) / 5))
	}
	return result
}

func (s PlainTextStatement) playFor(aPerformance Performance) Play {
	return s.plays[aPerformance.PlayID]
}

func (s PlainTextStatement) amountFor(aPerformance Performance) (int, error) {
	result := 0
	switch s.playFor(aPerformance).Type {
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
		return 0, fmt.Errorf("unknown type: %s", s.playFor(aPerformance).Type)
	}
	return result, nil
}

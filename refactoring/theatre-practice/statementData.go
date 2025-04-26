package theatre

import (
	"fmt"
	"math"
)

type StatementData struct {
	Customer           string
	Performances       []enrichedPerformance
	TotalAmount        int
	TotalVolumeCredits int
	plays              map[string]Play
}

type enrichedPerformance struct {
	Name     string
	Amount   int
	Audience int
	Type     string
}

func CreateStatementData(invoice Invoice, plays map[string]Play) (StatementData, error) {
	result := StatementData{}
	result.plays = plays
	result.Customer = invoice.Customer

	enrichedPerformances := []enrichedPerformance{}
	for _, perf := range invoice.Performances {
		ep, err := result.enrichPerformance(perf)
		if err != nil {
			return StatementData{}, err
		}
		enrichedPerformances = append(enrichedPerformances, ep)
	}
	result.Performances = enrichedPerformances

	result.TotalAmount = result.totalAmount()
	result.TotalVolumeCredits = result.totalVolumeCredits()

	return result, nil
}

func (s StatementData) enrichPerformance(perf Performance) (enrichedPerformance, error) {
	result := enrichedPerformance{}
	amount, err := s.amount(perf)
	if err != nil {
		return enrichedPerformance{}, err
	}
	result.Amount = amount
	result.Audience = perf.Audience
	result.Name = s.playFor(perf).Name
	result.Type = s.playFor(perf).Type

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

func (s StatementData) totalAmount() int {
	result := 0
	for _, perf := range s.Performances {
		result += perf.Amount
	}
	return result
}

func (s StatementData) totalVolumeCredits() int {
	result := 0
	for _, perf := range s.Performances {
		result += s.volumeCreditsFor(perf)
	}
	return result
}

func (StatementData) volumeCreditsFor(perf enrichedPerformance) int {
	result := 0
	result += int(math.Max(float64(perf.Audience)-30, 0))

	if perf.Type == "comedy" {
		result += int(math.Floor(float64(perf.Audience) / 5))
	}
	return result
}

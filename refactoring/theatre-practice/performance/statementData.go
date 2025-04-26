package performance

import (
	"dsa/refactoring/theatre-practice/common"
	"fmt"
	"math"
)

type StatementData struct {
	Customer           string
	Performances       []enrichedPerformance
	TotalAmount        int
	TotalVolumeCredits int
	plays              map[string]common.Play
}

type enrichedPerformance struct {
	playID   string
	Name     string
	Amount   int
	Audience int
	theType  string
}

func CreateStatementData(invoice common.Invoice, plays map[string]common.Play) (StatementData, error) {
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

func (s StatementData) enrichPerformance(perf common.Performance) (enrichedPerformance, error) {
	result := enrichedPerformance{}
	result.Audience = perf.Audience
	result.playID = perf.PlayID
	result.Name = s.playFor(result).Name
	result.theType = s.playFor(result).Type

	amount, err := s.amount(result)
	if err != nil {
		return enrichedPerformance{}, err
	}
	result.Amount = amount

	return result, nil
}

func (StatementData) amount(perf enrichedPerformance) (int, error) {
	result := 0
	switch perf.theType {
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
		return 0, fmt.Errorf("unknown type: %s", perf.theType)
	}
	return result, nil
}

func (s StatementData) playFor(perf enrichedPerformance) common.Play {
	return s.plays[perf.playID]
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

	if perf.theType == "comedy" {
		result += int(math.Floor(float64(perf.Audience) / 5))
	}
	return result
}

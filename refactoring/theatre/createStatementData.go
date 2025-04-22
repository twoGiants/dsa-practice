package theatre

import (
	"fmt"
	"math"
)

type StatementData struct {
	Customer           string
	Performances       []EnrichedPerformance
	TotalAmount        int
	TotalVolumeCredits int
}

type EnrichedPerformance struct {
	PlayID        string
	Audience      int
	play          Play
	amount        int
	volumeCredits int
}

type Plays map[string]Play

func (p Plays) playFor(aPerformance Performance) Play {
	return p[aPerformance.PlayID]
}

func createStatementData(invoice Invoice, plays map[string]Play) (StatementData, error) {
	result := StatementData{}
	result.Customer = invoice.Customer

	var enrichedPerformances []EnrichedPerformance
	for _, performance := range invoice.Performances {
		enrichedPerformance, err := enrichPerformance(performance, Plays(plays))
		if err != nil {
			return StatementData{}, err
		}
		enrichedPerformances = append(enrichedPerformances, enrichedPerformance)
	}
	result.Performances = enrichedPerformances

	result.TotalAmount = totalAmount(result)
	result.TotalVolumeCredits = totalVolumeCredits(result)

	return result, nil
}

type PerformanceCalculator struct {
	performance Performance
	play        Play
}

func NewPerformanceCalculator(pe Performance, pl Play) PerformanceCalculator {
	return PerformanceCalculator{pe, pl}
}

func (p PerformanceCalculator) Amount() (int, error) {
	result := 0
	switch p.play.Type {
	case "tragedy":
		result = 40000
		if p.performance.Audience > 30 {
			result += 1000 * (p.performance.Audience - 30)
		}
	case "comedy":
		result = 30000
		if p.performance.Audience > 20 {
			result += 10000 + 500*(p.performance.Audience-20)
		}
		result += 300 * p.performance.Audience
	default:
		return 0, fmt.Errorf("unknown type: %s", p.play.Type)
	}
	return result, nil
}

func (p PerformanceCalculator) VolumeCredits() int {
	result := int(math.Max(float64(p.performance.Audience)-30, 0))
	// add extra credit for every ten comedy attendees
	if p.play.Type == "comedy" {
		result += int(math.Floor(float64(p.performance.Audience) / 5))
	}
	return result
}

func enrichPerformance(performance Performance, plays Plays) (EnrichedPerformance, error) {
	calculator := NewPerformanceCalculator(performance, plays.playFor(performance))
	result := EnrichedPerformance{}
	result.PlayID = performance.PlayID
	result.Audience = performance.Audience
	result.play = calculator.play
	amount, err := calculator.Amount()
	if err != nil {
		return EnrichedPerformance{}, err
	}
	result.amount = amount
	result.volumeCredits = calculator.VolumeCredits()

	return result, nil
}

func totalVolumeCredits(data StatementData) int {
	result := 0
	for _, perf := range data.Performances {
		result += perf.volumeCredits
	}
	return result
}

func totalAmount(data StatementData) int {
	result := 0
	for _, perf := range data.Performances {
		result += perf.amount
	}
	return result
}

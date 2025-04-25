package performance

import (
	"dsa/refactoring/theatre/types"
	"fmt"
	"math"
)

type plays map[string]types.Play

func (p plays) playFor(aPerformance types.Performance) types.Play {
	return p[aPerformance.PlayID]
}

type enrichedPerformance struct {
	Audience      int
	Amount        int
	Play          types.Play
	playID        string
	volumeCredits int
}

type StatementData struct {
	Customer     string
	Performances []enrichedPerformance
}

func CreateStatementData(
	invoice types.Invoice,
	thePlays map[string]types.Play,
) (StatementData, error) {
	result := StatementData{}
	result.Customer = invoice.Customer

	var enrichedPerformances []enrichedPerformance
	for _, performance := range invoice.Performances {
		enrichedPerformance, err := result.enrichPerformance(performance, plays(thePlays))
		if err != nil {
			return StatementData{}, err
		}
		enrichedPerformances = append(enrichedPerformances, enrichedPerformance)
	}
	result.Performances = enrichedPerformances

	return result, nil
}

func (StatementData) enrichPerformance(
	performance types.Performance,
	plays plays,
) (enrichedPerformance, error) {
	calculator, err := createPerformanceCalculator(performance, plays.playFor(performance))
	if err != nil {
		return enrichedPerformance{}, err
	}

	result := enrichedPerformance{}
	result.playID = performance.PlayID
	result.Audience = performance.Audience
	result.Play = calculator.Play()
	result.Amount = calculator.Amount()
	result.volumeCredits = calculator.VolumeCredits()

	return result, nil
}

func (s StatementData) TotalVolumeCredits() int {
	result := 0
	for _, perf := range s.Performances {
		result += perf.volumeCredits
	}
	return result
}

func (s StatementData) TotalAmount() int {
	result := 0
	for _, perf := range s.Performances {
		result += perf.Amount
	}
	return result
}

type performanceCalculator interface {
	Amount() int
	VolumeCredits() int
	Play() types.Play
}

type performanceCalculatorImpl struct {
	performance types.Performance
	play        types.Play
}

func newPerformanceCalculator(pe types.Performance, pl types.Play) performanceCalculatorImpl {
	return performanceCalculatorImpl{pe, pl}
}

func createPerformanceCalculator(
	pe types.Performance,
	pl types.Play,
) (performanceCalculator, error) {
	switch pl.Type {
	case "tragedy":
		return tragedyCalculatorImpl{newPerformanceCalculator(pe, pl)}, nil
	case "comedy":
		return comedyCalculatorImpl{newPerformanceCalculator(pe, pl)}, nil
	default:
		return nil, fmt.Errorf("unknown type: %s", pl.Type)
	}
}

func (p performanceCalculatorImpl) Play() types.Play {
	return p.play
}

func (p performanceCalculatorImpl) VolumeCredits() int {
	return int(math.Max(float64(p.performance.Audience)-30, 0))
}

type tragedyCalculatorImpl struct {
	performanceCalculatorImpl
}

func (t tragedyCalculatorImpl) Amount() int {
	result := 40000
	if t.performance.Audience > 30 {
		result += 1000 * (t.performance.Audience - 30)
	}
	return result
}

type comedyCalculatorImpl struct {
	performanceCalculatorImpl
}

func (c comedyCalculatorImpl) Amount() int {
	result := 30000
	if c.performance.Audience > 20 {
		result += 10000 + 500*(c.performance.Audience-20)
	}
	result += 300 * c.performance.Audience
	return result
}

func (c comedyCalculatorImpl) VolumeCredits() int {
	return c.performanceCalculatorImpl.VolumeCredits() + int(
		math.Floor(float64(c.performance.Audience)/5),
	)
}

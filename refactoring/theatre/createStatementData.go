package theatre

import (
	"fmt"
	"math"
)

type Plays map[string]Play

func (p Plays) PlayFor(aPerformance Performance) Play {
	return p[aPerformance.PlayID]
}

type EnrichedPerformance struct {
	Audience      int
	Amount        int
	Play          Play
	playID        string
	volumeCredits int
}

type StatementData struct {
	Customer     string
	Performances []EnrichedPerformance
}

func CreateStatementData(invoice Invoice, plays map[string]Play) (StatementData, error) {
	result := StatementData{}
	result.Customer = invoice.Customer

	var enrichedPerformances []EnrichedPerformance
	for _, performance := range invoice.Performances {
		enrichedPerformance, err := result.enrichPerformance(performance, Plays(plays))
		if err != nil {
			return StatementData{}, err
		}
		enrichedPerformances = append(enrichedPerformances, enrichedPerformance)
	}
	result.Performances = enrichedPerformances

	return result, nil
}

func (StatementData) enrichPerformance(performance Performance, plays Plays) (EnrichedPerformance, error) {
	calculator, err := CreatePerformanceCalculator(performance, plays.PlayFor(performance))
	if err != nil {
		return EnrichedPerformance{}, err
	}

	result := EnrichedPerformance{}
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

type PerformanceCalculator interface {
	Amount() int
	VolumeCredits() int
	Play() Play
}

type PerformanceCalculatorImpl struct {
	performance Performance
	play        Play
}

func NewPerformanceCalculator(pe Performance, pl Play) PerformanceCalculatorImpl {
	return PerformanceCalculatorImpl{pe, pl}
}

func CreatePerformanceCalculator(pe Performance, pl Play) (PerformanceCalculator, error) {
	switch pl.Type {
	case "tragedy":
		return TragedyCalculatorImpl{NewPerformanceCalculator(pe, pl)}, nil
	case "comedy":
		return ComedyCalculatorImpl{NewPerformanceCalculator(pe, pl)}, nil
	default:
		return nil, fmt.Errorf("unknown type: %s", pl.Type)
	}
}

func (p PerformanceCalculatorImpl) Play() Play {
	return p.play
}

func (p PerformanceCalculatorImpl) VolumeCredits() int {
	return int(math.Max(float64(p.performance.Audience)-30, 0))
}

type TragedyCalculatorImpl struct {
	PerformanceCalculatorImpl
}

func (t TragedyCalculatorImpl) Amount() int {
	result := 40000
	if t.performance.Audience > 30 {
		result += 1000 * (t.performance.Audience - 30)
	}
	return result
}

type ComedyCalculatorImpl struct {
	PerformanceCalculatorImpl
}

func (c ComedyCalculatorImpl) Amount() int {
	result := 30000
	if c.performance.Audience > 20 {
		result += 10000 + 500*(c.performance.Audience-20)
	}
	result += 300 * c.performance.Audience
	return result
}

func (c ComedyCalculatorImpl) VolumeCredits() int {
	return c.PerformanceCalculatorImpl.VolumeCredits() + int(math.Floor(float64(c.performance.Audience)/5))
}

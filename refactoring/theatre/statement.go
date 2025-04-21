package theatre

import (
	"fmt"

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

type StatementPrinter struct{}

func (s StatementPrinter) Print(invoice Invoice, plays map[string]Play) (string, error) {
	statementData, err := createStatementData(invoice, plays)
	if err != nil {
		return "", err
	}

	return renderPlainText(statementData), nil
}

func renderPlainText(data StatementData) string {
	result := fmt.Sprintf("Statement for %s\n", data.Customer)

	for _, perf := range data.Performances {
		result += fmt.Sprintf(
			"  %s: %s (%d seats)\n",
			perf.play.Name,
			usd(perf.amount),
			perf.Audience,
		)
	}

	result += fmt.Sprintf("Amount owed is %s\n", usd(data.TotalAmount))
	result += fmt.Sprintf("You earned %d credits\n", data.TotalVolumeCredits)
	return result
}

func usd(number int) string {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	return ac.FormatMoney(number / 100)
}

package theatre

import (
	"fmt"

	"github.com/leekchan/accounting"
)

type StatementPrinter struct{}

func (s StatementPrinter) Print(invoice Invoice, plays map[string]Play) (string, error) {
	statementData, err := CreateStatementData(invoice, plays)
	if err != nil {
		return "", err
	}

	return renderPlayText(statementData)
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

func usd(amount int) string {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	return ac.FormatMoney(float64(amount) / 100)
}

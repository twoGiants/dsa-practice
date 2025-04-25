package theatre

import (
	perf "dsa/refactoring/theatre/performance"
	"dsa/refactoring/theatre/types"
	"fmt"

	"github.com/leekchan/accounting"
)

type StatementPrinter struct{}

func (s StatementPrinter) Print(
	invoice types.Invoice,
	plays map[string]types.Play,
) (string, error) {
	statementData, err := perf.CreateStatementData(invoice, plays)
	if err != nil {
		return "", err
	}

	return renderPlainText(statementData), nil
}

func renderPlainText(data perf.StatementData) string {
	result := fmt.Sprintf("Statement for %s\n", data.Customer)

	for _, perf := range data.Performances {
		result += fmt.Sprintf(
			"  %s: %s (%d seats)\n",
			perf.Play.Name,
			usd(perf.Amount),
			perf.Audience,
		)
	}

	result += fmt.Sprintf("Amount owed is %s\n", usd(data.TotalAmount()))
	result += fmt.Sprintf("You earned %d credits\n", data.TotalVolumeCredits())
	return result
}

func (s StatementPrinter) HtmlPrint(
	invoice types.Invoice,
	plays map[string]types.Play,
) (string, error) {
	statementData, err := perf.CreateStatementData(invoice, plays)
	if err != nil {
		return "", err
	}

	return renderHtml(statementData), nil
}

func renderHtml(data perf.StatementData) string {
	result := fmt.Sprintf("<h1>Statement for %s</h1>\n", data.Customer)
	result += "<table>\n"
	result += "<tr><th>play</th><th>seats</th><th>cost</th></tr>\n"
	for _, perf := range data.Performances {
		result += fmt.Sprintf(
			"  <tr><td>%s</td><td>%s</td><td>%d</td></tr>\n",
			perf.Play.Name,
			usd(perf.Amount),
			perf.Audience,
		)
	}
	result += "</table>\n"
	result += fmt.Sprintf("<p>Amount owed is <em>%s</em></p>\n", usd(data.TotalAmount()))
	result += fmt.Sprintf("<p>You earned <em>%d</em> credits</p>\n", data.TotalVolumeCredits())
	return result
}

func usd(number int) string {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	return ac.FormatMoney(number / 100)
}

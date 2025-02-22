package renderer

import (
	sch "example.com/amortization/schedule"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const border = "================================================================================"

var printer = message.NewPrinter(language.English)

func Format(sched sch.Schedule) string {
	result := ""
	result += border + "\n"
	result += printer.Sprintf("|%15s|%20s|%20s|%20s|\n", "Month", "Interest", "Paid to Principal", "Remaining Balance")
	result += border + "\n"
	for index, row := range sched.Rows {
		result += printer.Sprintf("|%15d|%20.2f|%20.2f|%20.2f|\n", index+1, row.PaidToInterest, row.PaidToPrincipal, row.RemainingBalance)
	}
	result += border + "\n"
	return result
}

func FormatMoney(value float64) string {
	return printer.Sprintf("%.2f", value)
}

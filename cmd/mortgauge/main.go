package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mdw-go/mortgauge"
)

func main() {
	config := parseConfig()
	iterator := mortgauge.NewAmortizationIterator(
		config.Principal,
		config.Interest,
		config.TermInMonths,
	)
	date := config.Start
	fmt.Println("Principal:", config.Principal)
	fmt.Println("Interest:", config.Interest)
	fmt.Println("TermInMonths:", config.TermInMonths)
	fmt.Println("Start:", date)
	fmt.Println("Monthly Payment:", iterator.MonthlyPayment())
	fmt.Printf("Date      Months    Balance   Principal Interest\n")
	fmt.Printf("================================================\n")
	for i := 0; iterator.NonZeroBalance(); i++ {
		step := iterator.Next(config.ExtraPayment)
		principal := fmt.Sprintf("%.2f", step.MonthlyPaymentOnPrincipal)
		interest := fmt.Sprintf("%.2f", step.MonthlyPaymentOnInterest)
		remaining := fmt.Sprintf("%.2f", step.RemainingPrincipal)
		fmt.Printf("%-9s %-9d %-9s %-9s %-9s\n",
			date.Format("2006-01"), config.TermInMonths-i, remaining, principal, interest,
		)
		date = date.AddDate(0, 1, 0)
	}
}

func parseConfig() (config Config) {
	flags := flag.NewFlagSet("mortgauge", flag.ContinueOnError)

	flags.Float64Var(&config.Principal, "principal", 100_000, "The original principal.")
	flags.Float64Var(&config.Interest, "interest", 6.0, "The interest rate, in percent.")
	flags.IntVar(&config.TermInMonths, "term", 180, "The term, in months. 180=15y; 360=30y;")
	flags.StringVar(&config.start, "start", "2020-01", "The month of the first payment ('YYYY-MM').")
	flags.Float64Var(&config.ExtraPayment, "extra", 0, "The extra principal to pay each month.")

	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	config.Start, err = time.Parse("2006-01", config.start)
	if err != nil {
		flags.PrintDefaults()
		log.Fatal(err)
	}

	return config
}

type Config struct {
	Principal    float64
	Interest     float64
	TermInMonths int
	Start        time.Time
	start        string
	ExtraPayment float64
}

// TODO: how to indicate extra payment (one-time, ongoing from <date>, etc...)

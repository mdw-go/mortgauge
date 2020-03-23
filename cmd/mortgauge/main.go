package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mdwhatcott/mortgauge"
)

func main() {
	config := parseConfig()
	iterator := mortgauge.NewAmortizationIterator(
		config.Principal,
		config.Interest,
		config.TermInMonths,
	)
	date := config.Start
	for i := 0; iterator.NonZeroBalance(); i++ {
		fmt.Println(date.Format("2006-01"), i+1, iterator.Next(config.ExtraPayment))
		date = date.AddDate(0, 1, 0)
	}
}

func parseConfig() (config Config) {
	flags := flag.NewFlagSet("mortgauge", flag.ContinueOnError)

	flags.Float64Var(&config.Principal, "principal", 100_000, "The original principal.")
	flags.Float64Var(&config.Interest, "interest", 6.0, "The interest rate, in percent.")
	flags.IntVar(&config.TermInMonths, "term", 180, "The term, in months.")
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

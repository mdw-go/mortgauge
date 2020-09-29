package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mdwhatcott/mortgauge"
)

func main() {
	config := parseConfig()
	totalExtra := 0.0
	extra := parseExtraPayments()
	iterator := mortgauge.NewAmortizationIterator(
		config.Principal,
		config.Interest,
		config.TermInMonths,
	)
	date := config.Start
	term := 0

	writer := csv.NewWriter(os.Stdout)
	writer.Comma = '\t'

	for ; iterator.NonZeroBalance(); term++ {
		extraPayment := extra[date]
		if !config.ExtraFrom.After(date) {
			extraPayment += config.Extra
		}
		state := iterator.Next(extraPayment)
		record := []string{
			fmt.Sprintf("%10.2f", state.StartingPrincipal),
			fmt.Sprintf("%10.2f", state.MonthlyPaymentOnInterest),
			fmt.Sprintf("%10.2f", state.MonthlyPaymentOnPrincipal-extraPayment),
			fmt.Sprintf("%10.2f", state.ExtraPaymentOnPrincipal),
			fmt.Sprintf("%10.2f", state.RemainingPrincipal),
			fmt.Sprint(term + 1),
			date.Format("2006-01"),
		}
		if date.Year() == time.Now().Year() && date.Month() == time.Now().Month() {
			record = append(record, "<<<")
		}
		_ = writer.Write(record)
		totalExtra += extraPayment
		date = date.AddDate(0, 1, 0)
	}

	_ = writer.Write([]string{
		fmt.Sprintf("%10s", "Starting"),
		fmt.Sprintf("%10s", "Interest"),
		fmt.Sprintf("%10s", "Principal"),
		fmt.Sprintf("%10s", "Extra"),
		fmt.Sprintf("%10s", "Remaining"),
		"Term",
		"Date",
	})

	writer.Flush()

	termDiff := config.TermInMonths - term - 1
	log.Printf(
		"Extra payments applied (%.2f) shortened term of loan by %d years and %d months, from %s to %s.",
		totalExtra,
		termDiff/12,
		termDiff%12,
		config.Start.AddDate(0, config.TermInMonths, 0).Format("2006-01"),
		config.Start.AddDate(0, term-1, 0).Format("2006-01"),
	)
}

func parseExtraPayments() map[time.Time]float64 {
	payments := make(map[time.Time]float64)
	reader := csv.NewReader(os.Stdin)
	records, err := reader.ReadAll()
	if err != nil {
		log.Panic(err)
	}
	for i, record := range records {
		if i == 0 { // header
			continue
		}
		rawDate := record[0]
		rawPrincipal := record[1]

		date, err := time.Parse("2006-01", rawDate)
		if err != nil {
			log.Panic(err)
		}
		principal, err := strconv.ParseFloat(rawPrincipal, 64)
		if err != nil {
			log.Panic(err)
		}
		payments[date] += principal
	}
	return payments
}

func parseConfig() (config Config) {
	flags := flag.NewFlagSet("mortgauge", flag.ContinueOnError)

	flags.Float64Var(&config.Principal, "principal", 100_000, "The original principal.")
	flags.Float64Var(&config.Interest, "interest", 6.0, "The interest rate, in percent.")
	flags.IntVar(&config.TermInMonths, "term", 180, "The term, in months.")
	flags.StringVar(&config.start, "start", "2020-01", "The month of the first payment ('YYYY-MM').")
	flags.Float64Var(&config.Extra, "extra", 0,
		"The extra payment on principle to apply each month, starting on -extra-from. "+
			"If left zero, no recurring extra payments are applied.")
	flags.StringVar(&config.extraFrom, "extra-from", "",
		"The month of the first recurring extra payment, specified by -extra. "+
			"If left blank, no recurring extra payments are applied.")

	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	config.Start, err = time.Parse("2006-01", config.start)
	if err != nil {
		flags.PrintDefaults()
		log.Fatal(err)
	}

	config.ExtraFrom, err = time.Parse("2006-01", config.extraFrom)
	if config.extraFrom != "" && err != nil {
		flags.PrintDefaults()
		log.Fatal(err)
	}

	return config
}

type Config struct {
	Principal    float64
	Interest     float64
	TermInMonths int

	Start time.Time
	start string

	Extra     float64
	ExtraFrom time.Time
	extraFrom string
}

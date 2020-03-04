package mortgauge

import (
	"math"
)

func CalculateMonthlyPayment(
	principal float64,
	interestPercent float64,
	termInMonths int,
) float64 {
	step1 := (interestPercent / 100.0) / 12
	step2 := math.Pow(1+step1, float64(termInMonths))
	step3 := step2 - 1
	step4 := step2 * step1
	step5 := step4 / step3
	return principal * step5
}

func AmortizationListing(
	principal float64,
	interestPercent float64,
	termInMonths int,
) (listing []Amortization) {
	payment := CalculateMonthlyPayment(principal, interestPercent, termInMonths)

	for principal > 0 {
		rate := (interestPercent / 100.0) / 12
		paymentOnInterest := principal * rate
		paymentOnPrincipal := payment - paymentOnInterest
		listing = append(listing, Amortization{
			StartingPrincipal:         principal,
			MonthlyPaymentOnPrincipal: paymentOnPrincipal,
			MonthlyPaymentOnInterest:  paymentOnInterest,
			RemainingPrincipal:        principal - paymentOnPrincipal,
		})

		principal -= paymentOnPrincipal
	}
	return listing
}

type Amortization struct {
	StartingPrincipal         float64
	MonthlyPaymentOnPrincipal float64
	ExtraPaymentOnPrincipal   float64
	MonthlyPaymentOnInterest  float64
	RemainingPrincipal        float64
}

type AmortizationIterator struct {
}

func NewAmortizationIterator(
	principal float64,
	interestPercent float64,
	termInMonths int,
) *AmortizationIterator {
	return &AmortizationIterator{}
}

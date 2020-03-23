package mortgauge

import "math"

func CalculateMonthlyPayment(
	principal,
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
	principal,
	interestPercent float64,
	termInMonths int,
) (listing []Amortization) {
	iterator := NewAmortizationIterator(principal, interestPercent, termInMonths)
	for iterator.NonZeroBalance() {
		listing = append(listing, iterator.Next(0))
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
	principal       float64
	interestPercent float64
	termInMonths    int
	payment         float64
}

func NewAmortizationIterator(
	principal,
	interestPercent float64,
	termInMonths int,
) *AmortizationIterator {
	payment := CalculateMonthlyPayment(principal, interestPercent, termInMonths)
	return &AmortizationIterator{
		principal:       principal,
		interestPercent: interestPercent,
		termInMonths:    termInMonths,
		payment:         payment,
	}
}

func (this *AmortizationIterator) NonZeroBalance() bool {
	return this.principal > 0
}

func (this *AmortizationIterator) Next(extraPayment float64) Amortization {
	rate := (this.interestPercent / 100.0) / 12
	paymentOnInterest := this.principal * rate
	paymentOnPrincipal := (this.payment - paymentOnInterest) + extraPayment
	defer this.applyPayment(paymentOnPrincipal)

	return Amortization{
		StartingPrincipal:         this.principal,
		MonthlyPaymentOnPrincipal: paymentOnPrincipal,
		MonthlyPaymentOnInterest:  paymentOnInterest,
		RemainingPrincipal:        this.principal - paymentOnPrincipal,
	}
}

func (this *AmortizationIterator) applyPayment(paymentOnPrincipal float64) {
	this.principal -= paymentOnPrincipal
}

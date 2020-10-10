package mortgauge

import (
	"reflect"
	"testing"
)

func TestMonthlyPayment(t *testing.T) {
	payment := CalculateMonthlyPayment(100000, 6.0, 240)
	assertEqual(t, 716.4310584781762, payment)
}

func TestAmortizationListing(t *testing.T) {
	listing := AmortizationListing(100000, 6.0, 360)
	if !assertEqual(t, 360, len(listing)) {
		return
	}
	expected := Amortization{
		StartingPrincipal:         100000,
		MonthlyPaymentOnPrincipal: 99.55052515275884,
		ExtraPaymentOnPrincipal:   0,
		MonthlyPaymentOnInterest:  500.00,
		RemainingPrincipal:        99900.44947484724,
	}
	assertEqual(t, expected, listing[0])
}

func assertEqual(t *testing.T, expected, actual interface{}) bool {
	if reflect.DeepEqual(expected, actual) {
		return true
	}
	t.Errorf("\nExpected: %+v\nActual:   %+v", expected, actual)
	return false
}

package mortgauge

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestCalculationFixture(t *testing.T) {
	gunit.Run(new(CalculationFixture), t)
}

type CalculationFixture struct {
	*gunit.Fixture
}

func (this *CalculationFixture) TestMonthlyPayment() {
	this.So(
		CalculateMonthlyPayment(100000, 6.0, 240),
		should.AlmostEqual, 716.43, .01,
	)
}

func (this *CalculationFixture) TestAmortizationListing() {
	listing := AmortizationListing(100000, 6.0, 360)
	this.So(listing, should.HaveLength, 360)
	this.So(listing[0].StartingPrincipal, should.Equal, 100000)
	this.So(listing[0].MonthlyPaymentOnPrincipal, should.AlmostEqual, 99.55, .01)
	this.So(listing[0].MonthlyPaymentOnInterest, should.AlmostEqual, 500.00, .01)
	this.So(listing[0].RemainingPrincipal, should.AlmostEqual, 99900.45, .01)
}

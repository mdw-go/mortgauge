package mortgauge

import (
	"testing"
	"time"

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

func (this *CalculationFixture) TestAmortizationListing2() {
	listing := AmortizationListing(165000, 3.0, 180)
	this.So(listing, should.HaveLength, 180)
	start := time.Date(2017, 6, 1, 0, 0, 0, 0, time.UTC)
	for i, item := range listing {
		this.Println(start.Format("2006-01"), i+1, item)
		start = start.AddDate(0, 1, 0)
	}
}

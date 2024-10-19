package mortgauge

import "testing"

func TestFormatUSD(t *testing.T) {
	assertEqual(t, "$0.00", FormatUSD(0))
	assertEqual(t, "$1.00", FormatUSD(1))
	assertEqual(t, "$10.00", FormatUSD(10))
	assertEqual(t, "$100.00", FormatUSD(100))
	assertEqual(t, "$1,000.00", FormatUSD(1_000))
	assertEqual(t, "$10,000.00", FormatUSD(10_000))
	assertEqual(t, "$100,000.00", FormatUSD(100_000))
	assertEqual(t, "$1,000,000.00", FormatUSD(1_000_000.00))

	assertEqual(t, "$-1.00", FormatUSD(-1))
	assertEqual(t, "$-10.00", FormatUSD(-10))
	assertEqual(t, "$-100.00", FormatUSD(-100))
	assertEqual(t, "$-1,000.00", FormatUSD(-1_000))
	assertEqual(t, "$-10,000.00", FormatUSD(-10_000))
	assertEqual(t, "$-100,000.00", FormatUSD(-100_000))
	assertEqual(t, "$-1,000,000.00", FormatUSD(-1_000_000.00))
}

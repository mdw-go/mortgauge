package mortgauge

import (
	"fmt"
	"strings"
)

func FormatUSD(usd float64) string {
	result := fmt.Sprintf("$%.2f", usd)
	decimal := strings.Index(result, ".")
	comma := decimal - 3
	edge := 1
	if usd < 0 { // account for negative sign
		edge = 2
	}
	for comma > edge {
		result = result[:comma] + "," + result[comma:]
		comma -= 3
	}
	return result
}

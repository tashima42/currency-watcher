package currencywatcher

import (
	"fmt"

	"github.com/tashima42/currency-watcher/pkg/currencyprovider"
)

func Check(currencyThreshold float64, returnAnyway bool) (*string, error) {
	conversionResult, err := currencyprovider.Convert(currencyprovider.USD, currencyprovider.BRL, 1)
	if err != nil {
		return nil, err
	}
	message := fmt.Sprintf("Threshold: %.2f\nConverted: %.2f\n", currencyThreshold, conversionResult.Converted)
	if conversionResult.Converted <= currencyThreshold {
		message += "TIME TO BUY"
		return &message, nil
	}
	if returnAnyway {
		message += "Not time to buy yet"
		return &message, nil
	}
	return nil, nil
}

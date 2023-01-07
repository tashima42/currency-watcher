package currencyconverter

import (
	"fmt"

	"github.com/tashima42/currency-watcher/pkg/currencyprovider"
)

func Convert(from string, to string, amount float64) (string, error) {
	conversionResult, err := currencyprovider.Convert(from, to, amount)
	fmt.Println(from, to, amount)
	if err != nil {
		return "", err
	}
	fmt.Println(conversionResult)
	return fmt.Sprintf("Converted: %.2f\n", conversionResult.Converted), nil
}

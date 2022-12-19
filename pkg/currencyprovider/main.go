package currencyprovider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	USD = "USD" // united states
	BRL = "BRL" // brazil
	CLP = "CLP" // chile
)

type ConversionResult struct {
	Base      string  `json:"base"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	Converted float64 `json:"converted"`
	Rate      float64 `json:"rate"`
	//LastUpdate time.Time `json:"lastUpdate"`
}

// Convert converts currency from one to another, it receives
// currencies in the first two parameters and the amount to be converted in the third
// and returns a string value as the result, or an error
func Convert(from string, to string, amount float64) (*ConversionResult, error) {
	baseUrl := os.Getenv("CURRENCY_EXCHANGE_BASEURL")
	apiKey := os.Getenv("CURRENCY_EXCHANGE_APIKEY")
	url := fmt.Sprintf("%s/exchange/convert?apiKey=%s&base=%s&to=%s&amount=%f", baseUrl, apiKey, from, to, amount)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var conversionResult ConversionResult
	err = json.NewDecoder(resp.Body).Decode(&conversionResult)
	if err != nil {
		return nil, err
	}

	return &conversionResult, nil
}

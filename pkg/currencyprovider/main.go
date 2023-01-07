package currencyprovider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	USD = "usd" // united states
	BRL = "brl" // brazil
	CLP = "clp" // chile
)

type ConversionResult struct {
	Converted float64 `json:"converted"`
	Rate      float64 `json:"rate"`
}

func Convert(from string, to string, amount float64) (*ConversionResult, error) {
	baseUrl := os.Getenv("CURRENCY_EXCHANGE_BASEURL")
	url := fmt.Sprintf("%s/latest/currencies/%s/%s.min.json", baseUrl, from, to)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var apiResult map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&apiResult)
	if err != nil {
		return nil, err
	}

	fmt.Println(apiResult)
	fmt.Println(apiResult[to])
	if err != nil {
		return nil, err
	}
	rate := apiResult[to].(float64)
	conversionResult := ConversionResult{
		Rate:      rate,
		Converted: rate * amount,
	}

	return &conversionResult, nil
}

// anyapi provider is disabled by now, it doesn't have CLP
/*
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
*/

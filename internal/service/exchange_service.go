package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeResponse struct {
	USDBRL struct {
		Bid  string `json:"bid"`
		Ask  string `json:"ask"`
		High string `json:"high"`
		Low  string `json:"low"`
	} `json:"USDBRL"`
}

func GetExchangeRate() (*ExchangeResponse, error) {
	resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch exchange rate: %w", err)
	}

	defer resp.Body.Close()

	var exchange ExchangeResponse
	if err := json.NewDecoder(resp.Body).Decode(&exchange); err != nil {
		return nil, fmt.Errorf("failed to decode exchange rate response: %w", err)
	}

	return &exchange, nil
}

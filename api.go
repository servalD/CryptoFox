package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"fmt"
)

// Fetch and parse kraken server Time
func getSystemTime() (SystemTimeResponse, error) {
	respTime, err := http.Get("https://api.kraken.com/0/public/Time")
	if err != nil {
		return SystemTimeResponse{}, fmt.Errorf("error fetching Time: %v", err)
	}
	defer respTime.Body.Close()

	body, err := io.ReadAll(respTime.Body)
	if err != nil {
		return SystemTimeResponse{}, fmt.Errorf("error reading Time: %v", err)
	}
	
	var value SystemTimeResponse
	err = json.Unmarshal(body, &value)
	if err != nil {
		return SystemTimeResponse{}, fmt.Errorf("error Unmarshaling Time: %v", err)
	}
	return value, nil
}

// Fetch and parse kraken server SystemStatus
func getSystemStatus() (SystemStatusResponse, error) {
	respStatus, err := http.Get("https://api.kraken.com/0/public/SystemStatus")
	if err != nil {
		return SystemStatusResponse{}, fmt.Errorf("error fetching SystemStatus: %v", err)
	}
	defer respStatus.Body.Close()
	
	body, err := io.ReadAll(respStatus.Body)
	if err != nil {
		return SystemStatusResponse{}, fmt.Errorf("error reading SystemStatus: %v", err)
	}
	
	var value SystemStatusResponse
	err = json.Unmarshal(body, &value)
	if err != nil {
		return SystemStatusResponse{}, fmt.Errorf("error unmarshaling systemStatus: %v", err)
	}
	return value, nil
}

// Fetch and parse all kraken server Tickers
func getTradingPairs() (TradingPairsResponse, error) {
	resp, err := http.Get("https://api.kraken.com/0/public/AssetPairs")
	if err != nil {
		return TradingPairsResponse{}, fmt.Errorf("error fetching AssetPairs: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TradingPairsResponse{}, fmt.Errorf("error reading AssetPairs: %v", err)
	}
	var value TradingPairsResponse
	err = json.Unmarshal(body, &value)
	if err != nil {
		return TradingPairsResponse{}, fmt.Errorf("error Unmarshaling AssetPairs: %v", err)
	}
	return value, nil
}

// Fetch and parse all/selected kraken server Tickers
func getTickerInfo(selected []string) (TickerInfoResponse, error) {
	var url string
	if len(selected) > 0 {
		url = fmt.Sprintf("https://api.kraken.com/0/public/Ticker?pair=%s", strings.Join(selected, ","))
	} else {
		url = "https://api.kraken.com/0/public/Ticker"
	}
	resp, err := http.Get(url)
	if err != nil {
		return TickerInfoResponse{}, fmt.Errorf("error fetching Ticker: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TickerInfoResponse{}, fmt.Errorf("error reading Ticker: %v", err)
	}
	var value TickerInfoResponse
	err = json.Unmarshal(body, &value)
	if err != nil {
		return TickerInfoResponse{}, fmt.Errorf("error Unmarshaling Ticker: %v", err)
	}
	return value, nil
}

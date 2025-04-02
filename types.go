package main

import (
	"log"
)

// Time 
type SystemTimeResponse struct {
	Error  []string `json:"error"`
	Result struct {
		Unixtime int    `json:"unixtime"`
		Rfc1123  string `json:"rfc1123"`
	} `json:"result"`
}

func (s *SystemTimeResponse) Print() {
	log.Printf("System Time:\n")
	log.Printf("Unix Time: %d\n", s.Result.Unixtime)
	log.Printf("RFC1123: %s\n", s.Result.Rfc1123)
}

// Status
type SystemStatusResponse struct {
	Error  []string `json:"error"`
	Result struct {
		Status string    `json:"status"`
	} `json:"result"`
}

func (s *SystemStatusResponse) Print() {
	log.Printf("System Status: %s\n", s.Result.Status)
}

// Trading Pairs
type TradingPairsResponse struct {
Error  []string `json:"error"`
	Result map[string]struct  {
			Wsname          string `json:"wsname"`
		}
}

func (t *TradingPairsResponse)GetFirst(count int) (list []string) {
	counter := 0
	for i := range t.Result {
		list = append(list, t.Result[i].Wsname)
		counter++
		if counter==count{
			break
		}
	}
	return
}

// Tickers
type TickerInfoResponse struct {
	Error  []string `json:"error"`
	Result map[string]struct  {
			LastTradePrice []string `json:"c"`
			High24h        []string `json:"h"`
			Low24h         []string `json:"l"`
		}
	}

func (t *TickerInfoResponse) Print() {
	log.Printf("Tickers Info:\n")
	for i := range t.Result {
		log.Printf("Ticker name: %s\n", i)
		log.Printf("Last Trade Price: %v\n", t.Result[i].LastTradePrice[0])
		log.Printf("High 24h: %v\n", t.Result[i].High24h[0])
		log.Printf("Low 24h: %v\n", t.Result[i].Low24h[0])
	}
	
}

// Tickers from db
// type Ticker struct{
// 	id 							 int
// 	last_trade_price string
// 	high_24h 				 string
// 	low_24h 				 string
// 	name 						 string
// 	local_unixtime 	 int
// 	server_unixtime  int
// 	server_status    string
// }

package main

func TrackKrackenPair(path string, pairCount int) (err error) {
	status, err := getSystemStatus()
	if err != nil {
		return 
	}
	
	time, err := getSystemTime()
	if err != nil {
		return 
	}

	pairs, err := getTradingPairs()
	if err != nil {
		return
	}
	
	// Select only first pairCount pairs for tickers
	tickers, err := getTickerInfo(pairs.GetFirst(pairCount))
	if err != nil {
		return
	}
	// Print the ticker info
	tickers.Print()
	// Insert the data into the database
	err = RegisterKrakenTickers(path, tickers, time, status)
	if err != nil {
		return
	}
	// err = KrakenTickersToCSV(path)
	return
}

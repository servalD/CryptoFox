package main

import (
	"log"
	"sync"
	_ "github.com/mattn/go-sqlite3" // Import sqlite3 driver
)

func main() {

	// channel := make(chan []Ticker)

	wg := sync.WaitGroup{}
	wg.Add(1)
	// Pair monitoring
	go func (){
		err := InitDatabase("./kraken_tickers.sqlite")
		if err != nil {
			log.Fatalf("Error initializing database: %v", err)
		}
		
		// TODO: Add a loop
		// Track the first 5 pairs
		err = TrackKrackenPair("./kraken_tickers.sqlite", 5)
		if err != nil {
			log.Fatal(err)
		}
		
		// channel <- []Ticker{}
		wg.Done()
	}()

	// API
	// go func (){
		
	// 	wg.Done()
	// }()
	
	wg.Wait()
}

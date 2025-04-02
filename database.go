package main

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	"time"
	_ "github.com/mattn/go-sqlite3" // Import sqlite3 driver
)

func InitDatabase(path string) error{
	// Remove DB file if it exists
	err := os.Remove(path)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error removing database file: %v", err)
	}

	// Open a new database connection
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Create a table for ticker info
	createTickerTableSQL := `
		CREATE TABLE IF NOT EXISTS ticker_info (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			last_trade_price TEXT,
			high_24h TEXT,
			low_24h TEXT,
			name TEXT,
			local_unixtime INTEGER,
			server_unixtime INTEGER,
			server_status TEXT
		);`
	_, err = db.Exec(createTickerTableSQL)
	if err != nil {
		return fmt.Errorf("error creating ticker info table: %v", err)
	}
	return nil
}

func RegisterKrakenTickers(path string, p TickerInfoResponse, t SystemTimeResponse, s SystemStatusResponse) error{

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	insertTickerSQL := `INSERT INTO ticker_info (last_trade_price, high_24h, low_24h, name, local_unixtime, server_unixtime, server_status) VALUES (?, ?, ?, ?, ?, ?, ?)`
	for i := range p.Result {
		_, err := db.Exec(insertTickerSQL, p.Result[i].LastTradePrice[0], p.Result[i].High24h[0], p.Result[i].Low24h[0], i, time.Now().Unix(), t.Result.Unixtime, s.Result.Status)
		if err != nil {
			log.Fatalf("error inserting ticker Tickers: %v", err)
		}
	}
	return nil
}

func GetRegisteredKrakenTickers(path string) ( error){
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return  fmt.Errorf("error opening database: %v", err)//[]Ticker{},
	}
	defer db.Close()

	var id 							 []int
	var last_trade_price []string
	var high_24h 				 []string
	var low_24h 				 []string
	var name 						 []string
	var local_unixtime 	 []int
	var server_unixtime  []int
	var server_status    []string

	err = db.QueryRow(`SELECT * from ticker_info`).Scan(&id, &last_trade_price, &high_24h, &low_24h, &name, &local_unixtime, &server_unixtime, &server_status)
	if err != nil {
		log.Fatalf("error getting Tickers: %v", err)
	}
	log.Printf("from db: %v, %v, %v, %v, %v, %v, %v, %v", id, last_trade_price, high_24h, low_24h, name, local_unixtime, server_unixtime, server_status)
	return nil
}

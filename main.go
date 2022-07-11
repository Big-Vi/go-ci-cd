package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"database/sql"

	"github.com/big-vi/go-ci-cd/config"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Print("failed to initialize the db", err)
	}
	defer db.Conn.Close()

	http.HandleFunc("/", HomePage(db.Conn))
	log.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomePage(conn *sql.DB) (http.HandlerFunc) {
	return func (w http.ResponseWriter, r *http.Request) {
		record, err := countRecords(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Records count:", record)

		type Record struct{
			Row int `json:row`
		}
		recordJSON := Record{Row: record}
		
		byteArray, err := json.Marshal(recordJSON)
		if err != nil {
			fmt.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(string(byteArray)))
	}
}

func countRecords(conn *sql.DB) (int, error) {
	rows, err := conn.Query("SELECT COUNT(*) FROM persons")
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	count := 0
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
		rows.Close()
	} 
	return count, nil
}


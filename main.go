package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"


	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func main() {
	db, err := initDB()
	if err != nil {
		log.Print("failed to initialize the db", err)
	}
	defer db.Conn.Close()

	http.HandleFunc("/", homePage(db.Conn))
	log.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

type Database struct {
	Conn *sql.DB
}

func initDB() (Database, error) {
	db := Database{}

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Creating connection string for DB
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
	)

	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		return db, err
	}

	db.Conn = conn

	if _, err := db.Conn.Exec("CREATE TABLE IF NOT EXISTS persons (PersonID int PRIMARY KEY,LastName varchar(255))"); err != nil {
		fmt.Println(err)
	}

	if _, err := db.Conn.Exec("INSERT INTO persons(PersonID, LastName) VALUES($1, $2) ON CONFLICT (PersonID) DO NOTHING", 1, "murugan"); err != nil {
		fmt.Println(err)
	}

	log.Println("DB connected")
	return db, nil
}

func homePage(conn *sql.DB) (http.HandlerFunc) {
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


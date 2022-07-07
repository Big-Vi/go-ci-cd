package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is home page.")
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Print("failed to initialize the db", err)
	}
	defer db.Conn.Close()

	http.HandleFunc("/", homePage)
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

	defer conn.Close()

	db.Conn = conn

	if _, err := db.Conn.Exec("CREATE TABLE IF NOT EXISTS persons (PersonID int PRIMARY KEY,LastName varchar(255))"); err != nil {
		fmt.Println(err)
	}

	rows, err := db.Conn.Query("SELECT COUNT(*) FROM persons")
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
	fmt.Println(count)

	log.Println("DB connected")
	return db, nil
}



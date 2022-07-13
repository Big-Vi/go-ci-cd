package config

import(
	"database/sql"
	"os"
	"log"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func InitDB() (Database, error) {
	db := Database{}

	// load .env file
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

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
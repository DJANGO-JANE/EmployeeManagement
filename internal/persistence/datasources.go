package persistence

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //postgres drivers for initialisation
	log "github.com/sirupsen/logrus"
)

type dataSources struct {
	DB *sqlx.DB
}

//Create database connection with Postgres
func CreateConnection() (*dataSources, error) {
	//load .env

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("ERROR : Error when loading .env")
	}

	//Retrieve connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DB"), os.Getenv("PG_SSL"))

	//Open connection
	log.Info("Attempting to connect using sqlx with connection string")
	db, cError := sqlx.Connect("postgres", connStr)

	if cError != nil {
		//panic(cError)

		fmt.Errorf("ERROR : Failed to open database connection : %w", cError)
		log.Fatal("ERROR : Failed to open database connection. ")
		fmt.Printf("The error is %s", cError)
	}

	//If ping fails, then database connection likely failed
	if err := db.Ping(); err != nil {
		log.Fatal("ERROR : Failed to connect to database")
		return nil, fmt.Errorf("ERROR : Error connecting to database. %w", err)
	}
	log.Info("Connection has been initialised.")
	return &dataSources{db}, nil
}

// An extension for this type for closing database connection
func (d *dataSources) Close() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("ERROR : Error closing database connection: %w", err)
	}
	return nil
}

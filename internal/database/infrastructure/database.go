package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"stori/internal/config"

	_ "github.com/lib/pq"
)

type PostgresConnector struct{}

func (p *PostgresConnector) Connect() (*sql.DB, error) {
	log.Println("Trying to connect to postgres")

	log.Printf("OS ENV VAriables: %+v", os.Environ())

	connectionDetails := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("STORI_POSTGRES_HOST"),
		os.Getenv("STORI_POSTGRES_PORT"),
		os.Getenv("STORI_POSTGRES_USER"),
		os.Getenv("STORI_POSTGRES_PASSWORD"),
		os.Getenv("STORI_POSTGRES_DBNAME"),
	)

	log.Println(connectionDetails)

	db, err := sql.Open("postgres", connectionDetails)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database is connected!")

	return db, nil
}

func (p *PostgresConnector) Disconnect() error {
	return config.Config().Database.Close()
}

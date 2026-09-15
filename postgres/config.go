package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sethvargo/go-envconfig"
)

var DB *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{}
	if err := envconfig.Process(context.Background(), &config); err != nil {
		log.Fatal(err)
	}

	log.Println("Connect to database")
	DB = ConnectToDb(&config)
}

type Config struct {
	User     string `env:"USER"`
	Password string `env:"PASS"`
	DbName   string `env:"DB_NAME"`
}

func ConnectToDb(config *Config) *sql.DB {
	connStr := "user=user password=1234 dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	return db
}

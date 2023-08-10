// cmd/api/main.go
package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type Config struct {
	Dsn     string `required:"true"`
	GinPort string `required:"true" split_words:"true"`
}

func LoadConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &c, nil
}

func NewPostgresDB(conf *Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", conf.Dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	conf, err := LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("configurations: ", conf)
	db, err := NewPostgresDB(conf)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("connected to db with: ", db)
}

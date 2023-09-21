package database

import (
	"context"
	"log"

	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/migrate"
)

var DBConn *ent.Client

func New() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=secret-password sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	DBConn = client

	return client
}

func DevelopmentMigrate() {
	if err := DBConn.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

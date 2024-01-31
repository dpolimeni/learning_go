package common

import (
	"fmt"
	"os"

	"github.com/dpolimeni/fiber_app/ent"
)

func GetDB() *ent.Client {
	err := LoadEnv()
	if err != nil {
		panic(err)
	}

	const (
		host   = "localhost"
		port   = 5432
		user   = "postgres"
		dbname = "gotest"
	)
	var password = os.Getenv("password")

	// Initialize the client
	connection := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	client, err := ent.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	return client

}

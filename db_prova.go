package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dpolimeni/fiber_app/common"
	"github.com/dpolimeni/fiber_app/ent"

	_ "github.com/lib/pq"
)

func main() {
	err := common.LoadEnv()
	password := os.Getenv("password")
	connection := fmt.Sprintf("host=localhost port=5432 user=postgres dbname=gotest password=%s sslmode=disable", password)
	fmt.Println(connection)
	client, err := ent.Open("postgres", connection)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"my-app-api/db"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	var now time.Time
	err := db.Pool.QueryRow(
		context.Background(),
		"SELECT now()",
	).Scan(&now)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB OK:", now)
}

package main

import (
	"context"

	"log"

	"music-library/internal/app"
)

// docker compose --env-file local.env up -d
// swag init -g cmd/main.go
func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = a.Run()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

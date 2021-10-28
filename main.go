package main

import (
	"context"
	"os"

	"github.com/alexdang1993374/short-link-engine/config"
	"github.com/alexdang1993374/short-link-engine/controllers"
)

func main() {
	db := config.Connect()
	ctx := context.Background()

	// Only creates table if it doesn't exist already
	controllers.CreateUrlTable(db, ctx)

	controllers.InsertUrl(db, ctx, os.Args[1])
}

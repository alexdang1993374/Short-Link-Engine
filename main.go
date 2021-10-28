package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alexdang1993374/short-link-engine/config"
	"github.com/alexdang1993374/short-link-engine/controllers"
)

func main() {
	db := config.Connect()
	ctx := context.Background()

	// Only creates table if it doesn't exist already
	controllers.CreateUrlTable(db, ctx)

	originalUrl := controllers.GetOriginalUrl(db, ctx, os.Args[1])

	shortUrl := controllers.GetShortUrl(db, ctx, os.Args[1])

	if originalUrl != "Not Found" {
		fmt.Println("original url:", originalUrl)
	} else if shortUrl != "Not Found" {
		fmt.Println("shortened url:", shortUrl)
	} else {
		fmt.Println("shortened url:", controllers.InsertUrl(db, ctx, os.Args[1]))
	}
}

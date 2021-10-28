package main

import (
	"os"

	"github.com/alexdang1993374/short-link-engine/controllers"
)

func main() {
	controllers.CreateUrlTable()
	controllers.InsertUrl(os.Args[1])
}

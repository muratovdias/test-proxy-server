package main

import (
	"log"

	"github.com/muratovdias/test-proxy-server/internal/pkg/app"
)

func main() {
	app := app.NewApp()
	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

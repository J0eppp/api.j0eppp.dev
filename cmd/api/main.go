package main

import (
	"github.com/J0eppp/api.j0eppp.dev/pkg/application"
	"github.com/J0eppp/api.j0eppp.dev/pkg/exithandler"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load env vars")
		log.Fatal(err.Error())
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("main() 1")

	exithandler.Init(func() {
		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}

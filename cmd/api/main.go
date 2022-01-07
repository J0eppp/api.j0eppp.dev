package main

import (
	"github.com/J0eppp/api.j0eppp.dev/pkg/application"
	"github.com/J0eppp/api.j0eppp.dev/pkg/exithandler"
	"github.com/J0eppp/api.j0eppp.dev/pkg/router"
	"github.com/J0eppp/api.j0eppp.dev/pkg/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Set the log output so we can see it in the docker console
	log.SetOutput(os.Stdout)

	// Get environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load env vars")
		log.Fatal(err.Error())
	}

	// Get the application
	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Get the server and setup default things
	srv := server.Get().WithAddress(app.Config.GetAPIPort()).WithRouter(router.Get(app)).WithErrorLogger(log.Default())

	// Start the server in a separate thread
	go func() {
		log.Printf("Starting server at %s\n", app.Config.GetAPIPort())
		if err := srv.Start(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	exithandler.Init(func() {
		if err := srv.Close(); err != nil {
			log.Println(err.Error())
		}

		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}

package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/derrickyoo/femProject/internal/app"
	"github.com/derrickyoo/femProject/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	chiMux := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      chiMux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Server is running on port %d", port)
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

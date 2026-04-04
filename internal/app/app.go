// Package app provides initialization and shared resources.
package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/derrickyoo/femProject/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// our stores will go here

	// our handlers will go here
	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}

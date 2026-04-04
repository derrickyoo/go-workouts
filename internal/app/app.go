// Package app provides initialization and shared resources.
package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/derrickyoo/go-workouts/internal/api"
	"github.com/derrickyoo/go-workouts/internal/store"
	"github.com/derrickyoo/go-workouts/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// our stores will go here
	workoutStore := store.NewPostgresWorkoutStore(pgDB)

	// our handlers will go here
	workoutHandler := api.NewWorkoutHandler(workoutStore)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}

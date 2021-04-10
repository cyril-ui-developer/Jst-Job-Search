package main

import (
	"fmt"
	"net/http"

	transportHTTP "github/cyril-ui-developer/JstJobSearch/internal/transport/http"
	"github/cyril-ui-developer/JstJobSearch/internal/db"
)

//App - the struct which contains things like pointers to db connections
type App struct{}

//Postgress db - docker run --name jobs-search-api-db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
func (app *App) Run() error {
	fmt.Println("Setting Our APP")
	
	var err error
	_, err = db.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("JST Job Saerch REST API, implemented in Golang")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up JST Job Saerch REST API")
		fmt.Println(err)
	}
}

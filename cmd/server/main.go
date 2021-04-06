package main

import (
	"fmt"
	transportHTTP "github.com/cyril-ui-developer/Jst-Job-Search/internal/transport/http"
)

//App - the struct which contains things like pointers to db connections
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting Our APP")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAnd.Serve(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set server")
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

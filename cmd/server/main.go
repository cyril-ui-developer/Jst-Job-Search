package main

import (
	"fmt"
	"net/http"

	"github/cyril-ui-developer/JstJobSearch/internal/db"
	"github/cyril-ui-developer/JstJobSearch/internal/jobs"
	transportHTTP "github/cyril-ui-developer/JstJobSearch/internal/transport/http"
	//"github/cyril-ui-developer/JstJobSearch/internal/migration"
	log "github.com/sirupsen/logrus"
)

//App - the struct which contains things like pointers to db connections
type App struct{
	Name string
    Version string
}

//Postgress db - docker run --name jobs-search-api-db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
// kind create cluster
// kubectl cluster-info --context kind-kind
// envsubst < config/deployment.yaml > temp.yml
// docker build -t forbsey/comment-api .
// docker push forsey/comment-api:latest
// envsubst < config/deployment.yaml | kubectl apply -f -
// kubectl apply -f config/service.yml
//kubectl port-forward service/comments-api 8080:8080

func (app *App) Run() error {
	log.Info("Setting Our APP")
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting Up Our APP")

	var err error
	database, err := db.NewDatabase()
	if err != nil {
		return err
	}

	err = db.MigrateDB(database)
	if err != nil {
		return err
	}

	jobService := jobs.NewService(database)

	handler := transportHTTP.NewHandler(jobService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	// =======
	// 	transportHTTP "github.com/cyril-ui-developer/Jst-Job-Search/internal/transport/http"
	// )

	// //App - the struct which contains things like pointers to db connections
	// type App struct{}

	// func (app *App) Run() error {
	// 	fmt.Println("Setting Our APP")

	// 	handler := transportHTTP.NewHandler()
	// 	handler.SetupRoutes()

	// 	if err := http.ListenAnd.Serve(":8080", handler.Router); err != nil {
	// 		fmt.Println("Failed to set server")
	// 		return err
	// 	}
	// >>>>>>> a9d7ec522e030f1064b6e1d60bc491d4f70cf2e1
	return nil
}

func main() {
	fmt.Println("JST Job Saerch REST API, implemented in Golang")

	app := App{
		Name: "JST Job API",
		Version: "1.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting up JST Job Saerch REST API")
		log.Fatal(err)
	}
}

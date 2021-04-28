module github/cyril-ui-developer/JstJobSearch

go 1.14

require (
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/sirupsen/logrus v1.8.1 // indirect
)

replace github.com/cyril-ui-developer/jst-job-search/internal/db => ../db

replace github.com/cyril-ui-developer/jst-job-search/internal/jobs => ../jobs

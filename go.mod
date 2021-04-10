module github/cyril-ui-developer/JstJobSearch

go 1.14

require (
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
)

replace github.com/cyril-ui-developer/jst-job-search/internal/db => ../db

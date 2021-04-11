package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github/cyril-ui-developer/JstJobSearch/internal/jobs"
)
//Handler - stores the pointer to the jobs service
type Handler struct{
	Router *mux.Router
	Service *jobs.Service
}

// Handler - returns a pointer to a Handler
func NewHandler(service *jobs.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

//SetupRoutes - sets up all the routes for the app
func (h *Handler) SetupRoutes(){
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "I'm alive!")
	})
	h.Router.HandleFunc("/api/jobs/{id}", h.GetJob).Methods("GET")
	h.Router.HandleFunc("/api/jobs",h.GetAllJobs).Methods("GET")
	h.Router.HandleFunc("/api/job",h.PostJob).Methods("POST")
	h.Router.HandleFunc("/api/jobs/{id}", h.UpdateJob).Methods("PUT")
	h.Router.HandleFunc("/api/jobs/{id}", h.DeleteJob).Methods("DELETE")
}
// GetJob
func (h *Handler) GetJob(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	
	jobID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse uint from ID")
	}
	job, err := h.Service.GetJob(uint(jobID))
	if err != nil {
		fmt.Fprintf(w, "Error retrieving job by Id")
	}
	fmt.Fprintf(w, "%+v", job)
}

// GetAllJobs
func (h *Handler) GetAllJobs(w http.ResponseWriter, r *http.Request){
	jobs, err := h.Service.GetAllJobs()

	if err != nil {
		fmt.Fprintf(w, "Error retrieving all jobs")
	}
	fmt.Fprintf(w, "%+v", jobs)
}
// PostJob
func (h *Handler) PostJob(w http.ResponseWriter, r *http.Request){
	job, err := h.Service.PostJob(jobs.Job{
		Slug:"/",
	})

	if err != nil {
		fmt.Fprintf(w, "Error posting job")
	}
	fmt.Fprintf(w, "%+v", job)
}

// UpdateJob
func (h *Handler) UpdateJob(w http.ResponseWriter, r *http.Request){
	job, err := h.Service.UpdateJob(2, jobs.Job{
		Slug:"/new",
	})
	// job, err := h.Service.UpdateJob(1, jobs.Job{
	// 	Slug:"/",
	// })

	if err != nil {
		fmt.Fprintf(w, "Error: Falied to update job")
	}
	fmt.Fprintf(w, "%+v", job)
}

//DeleteJob -
func (h *Handler) DeleteJob(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	jobID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse uint from ID")
	}

	err = h.Service.DeleteJob(uint(jobID))
	if err != nil {
		fmt.Fprintf(w, "Error: Falied to delete job")
	}

	fmt.Fprintf(w, "Job successfully deleted")
}

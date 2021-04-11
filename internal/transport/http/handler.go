package http

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/gorilla/mux"
	"github/cyril-ui-developer/JstJobSearch/internal/jobs"
)

// import (
// 	"fmt"
// 	"github.com/gorilla/mux"

// )
//Handler - stores the pointer to the jobs service
type Handler struct{
	Router *mux.Router
	Service *jobs.Service
}

// Response - an object to store error
type Response struct {
	Message string
	Error string
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I'm alive"}); err != nil {
			panic(err)
		}
	})
	h.Router.HandleFunc("/api/jobs/{id}", h.GetJob).Methods("GET")
	h.Router.HandleFunc("/api/jobs",h.GetAllJobs).Methods("GET")
	h.Router.HandleFunc("/api/job",h.PostJob).Methods("POST")
	h.Router.HandleFunc("/api/jobs/{id}", h.UpdateJob).Methods("PUT")
	h.Router.HandleFunc("/api/jobs/{id}", h.DeleteJob).Methods("DELETE")
}
// GetJob
func (h *Handler) GetJob(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	
	jobID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(w, "Unable to parse uint from ID", err)
	}
	job, err := h.Service.GetJob(uint(jobID))
	if err != nil {
		errorResponse(w, "Error retrieving job by Id", err)
	}
	if err := json.NewEncoder(w).Encode(job); err != nil {
		panic(err)
	}
}

// GetAllJobs
func (h *Handler) GetAllJobs(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	jobs, err := h.Service.GetAllJobs()

	if err != nil {
		errorResponse(w, "Error retrieving all jobs", err)
	}
	if err := json.NewEncoder(w).Encode(jobs); err != nil {
		panic(err)
	}
}
// PostJob
func (h *Handler) PostJob(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)

	var job jobs.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
	   // json.NewEncoder(w).Encode(Response{Message: "Failed to decode JSON body"})
		errorResponse(w, "Failed to decodde JSON body", err)
		return
	}

	job, err := h.Service.PostJob(job)

	if err != nil {
		fmt.Fprintf(w, "Error posting job")
	}

	if err := json.NewEncoder(w).Encode(job); err != nil {
		panic(err)
	}
}

// UpdateJob
func (h *Handler) UpdateJob(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var updateJob jobs.Job 
	if err := json.NewDecoder(r.Body).Decode(&updateJob); err != nil {
		// json.NewEncoder(w).Encode(Response{Message: "Failed to decode JSON body"})
		// return
		errorResponse(w, "Failed to decode JSON body", err)
	}

	vars := mux.Vars(r)
	id := vars["id"]
	jobID, err := strconv.ParseUint(id, 10, 64)

	job, err := h.Service.UpdateJob(uint(jobID), updateJob)

	if err != nil {
		errorResponse(w, "Error: Falied to update job", err)
	}
	
	if err := json.NewEncoder(w).Encode(job); err != nil {
		panic(err)
	}
}

//DeleteJob -
func (h *Handler) DeleteJob(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	jobID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(w, "Unable to parse uint from ID", err)
	}

	err = h.Service.DeleteJob(uint(jobID))
	if err != nil {
		errorResponse(w, "Error: Falied to delete job", err)
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Job Successfully Deleted."}); err != nil {
		panic(err)
	}
}
// =======
// }

// func NewHandler() *Handler {
// 	return &Handler{}
// }

// //SetupRoutes - sets up all the routes for the app
// func (h *Haandler) SetupRoutes(){
// 	fmt.Println("Setting Up Routes")
// 	h.Router = mux.NewHandler()
// 	h.Router.HaandleFunc("/api/health", funct(w http.ResponseWriter, *http.Request){
// 		fmt.Println("w, I'm alive!")
// 	})
// }

func errorResponse(w http.ResponseWriter, message string, err error){
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}


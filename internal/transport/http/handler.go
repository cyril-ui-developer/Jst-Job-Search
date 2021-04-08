package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)
//Handler - stores the pointer to the jobs service
type Handler struct{
	Router *mux.Router
}

// Handler - returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

//SetupRoutes - sets up all the routes for the app
func (h *Handler) SetupRoutes(){
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "I'm alive!")
	})
}
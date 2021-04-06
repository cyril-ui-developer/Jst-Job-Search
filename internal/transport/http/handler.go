package github.com/cyril-ui-developer/Jst-Job-Search/internal/transport/http

import (
	"fmt"
	"github.com/gorilla/mux"
)
//Handler - stores the pointer to the jobs service
type Handler struct{
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

//SetupRoutes - sets up all the routes for the app
func (h *Haandler) SetupRoutes(){
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewHandler()
	h.Router.HaandleFunc("/api/health", funct(w http.ResponseWriter, *http.Request){
		fmt.Println("w, I'm alive!")
	})
}
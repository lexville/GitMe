package router

import (
	"GitMe/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes contains the routes that this application
// will uses
func Routes() {
	r := mux.NewRouter()
	repoController := controllers.AddViewTemplates()
	r.HandleFunc("/", repoController.HomeHandler).Methods("GET")
	r.HandleFunc("/user", nil).Methods("POST")
	r.HandleFunc("/user/{username}", nil).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":3000", r)
}

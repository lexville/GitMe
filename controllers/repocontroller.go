package controllers

import (
	"GitMe/view"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Repocontroller contains all the fields needed by
// the repoctroller
type Repocontroller struct {
	HomeView *view.AppView
}

// AddViewTemplates adds all the templates
// needed by the repocontroller
func AddViewTemplates() *Repocontroller {
	return &Repocontroller{
		HomeView: view.AddTempateFiles(
			"base",
			"templates/repo/repo.gohtml",
		),
	}
}

// HomeHandler is responsible for the home view
//
// GET: /
func (rc *Repocontroller) HomeHandler(w http.ResponseWriter, r *http.Request) {
	rc.HomeView.Render(w, nil)
}

// PostUserHandler is responsible for the home view
//
// POST: /
func (rc *Repocontroller) PostUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	http.Redirect(w, r, "/user/ "+username, http.StatusSeeOther)
}

// GetUserHandler is responsible for the home view
//
// GET: /user/:username
func (rc *Repocontroller) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	fmt.Fprint(w, username)
}

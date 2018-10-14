package controllers

import (
	"GitMe/view"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Repocontroller contains all the fields needed by
// the repoctroller
type Repocontroller struct {
	HomeView *view.AppView
	RepoView *view.AppView
}

// AddViewTemplates adds all the templates
// needed by the repocontroller
func AddViewTemplates() *Repocontroller {
	return &Repocontroller{
		HomeView: view.AddTempateFiles(
			"base",
			"templates/repo/repo.gohtml",
		),
		RepoView: view.AddTempateFiles(
			"base",
			"templates/repo/user-repo.gohtml",
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
	http.Redirect(w, r, "/user/"+username, http.StatusSeeOther)
}

type UserNotFound struct {
	ErrorType string
	Message   string
}

// GetUserHandler is responsible for the home view
//
// GET: /user/:username
func (rc *Repocontroller) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	response := apiCall(username)
	defer response.Body.Close()
	if response.StatusCode == 404 {
		UserNotFound := UserNotFound{
			ErrorType: "failure",
			Message:   "No user in github has that username",
		}
		rc.RepoView.Render(w, UserNotFound)
		return
	}
	fmt.Fprint(w, username)
}

func apiCall(username string) *http.Response {
	response, err := http.Get("https://api.github.com/users/" + username)
	if err != nil {
		log.Fatal("Unable to make the request: ", err)
	}
	return response
}

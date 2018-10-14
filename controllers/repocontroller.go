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
			"templates/repo/not-found-user.gohtml",
		),
	}
}

// HomeHandler is responsible for the home view
//
// GET: /
func (rc *Repocontroller) HomeHandler(w http.ResponseWriter, r *http.Request) {
	rc.HomeView.Render(w, nil)
}

// SearchUserHandler is responsible for searching
// a user
//
// GET: /user
func (rc *Repocontroller) SearchUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	http.Redirect(w, r, "/user/"+username, http.StatusSeeOther)
}

// UserNotFound contains the error type
// and message for when a user cannot be found
type UserNotFound struct {
	ErrorType string
	Message   string
}

// GetUserHandler is responsible for getting the user data
//
// GET: /user/:username
func (rc *Repocontroller) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	getUserURL := "https://api.github.com/users/" + username
	response := apiCall(getUserURL, username)
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

func apiCall(theURL string, username string) *http.Response {
	response, err := http.Get(theURL)
	if err != nil {
		log.Fatal("Unable to make the request: ", err)
	}
	return response
}

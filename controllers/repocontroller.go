package controllers

import (
	"GitMe/view"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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
			"templates/repo/repo-user-data.gohtml",
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

// UserData contains all the data associated with the
// github user
type UserData struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           string      `json:"company"`
	Blog              string      `json:"blog"`
	Location          interface{} `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               interface{} `json:"bio"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

// User contains the user data if the user ezists and
// the error type and message if the user doesn't exist
type User struct {
	UserData
	UserNotFound
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
		userNotFound := UserNotFound{
			ErrorType: "failure",
			Message:   "No user in github has that username",
		}
		user := UserData{}
		data := User{
			user,
			userNotFound,
		}
		rc.RepoView.Render(w, data)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error while reading the body: ", err)
	}
	userData := UserData{}
	err = json.Unmarshal(body, &userData)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}
	userNotFound := UserNotFound{}
	data := User{
		userData,
		userNotFound,
	}
	rc.RepoView.Render(w, data)
	// fmt.Fprintln(w, data)
}

func apiCall(theURL string, username string) *http.Response {
	response, err := http.Get(theURL)
	if err != nil {
		log.Fatal("Unable to make the request: ", err)
	}
	return response
}

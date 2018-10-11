package controllers

import (
	"GitMe/view"
	"net/http"
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

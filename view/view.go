package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// AppView contains a template as well as the
// default layout to be used
type AppView struct {
	Template *template.Template
	Layout   string
}

// AddTempateFiles parses the template files needed
// to generate a specific view
func AddTempateFiles(layout string, files ...string) *AppView {
	files = append(
		files,
		getLayoutFiles()...,
	)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal("Unable to parse the template files")
	}
	return &AppView{
		Template: t,
		Layout:   layout,
	}
}

var (
	layoutDirectory = "templates/layout/"
	layoutExtension = ".gohtml"
)

func getLayoutFiles() []string {
	files, err := filepath.Glob(layoutDirectory + "*" + layoutExtension)
	if err != nil {
		log.Fatal("Unable to get the layout files: ", err)
	}
	return files
}

// Render renders the view needed
func (v *AppView) Render(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	if err := v.Template.ExecuteTemplate(w, v.Layout, data); err != nil {
		log.Fatal("Unable to render a template view: ", err)
	}
}

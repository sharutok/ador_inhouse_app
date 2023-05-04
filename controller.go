package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
}

func (app *application) AppLink(res http.ResponseWriter, req *http.Request) {

	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")

	if err != nil {
		app.logger.Println(err)
	}
	res.WriteHeader(http.StatusOK)
	res.Write(js)
}

func (app *application) linkage(res http.ResponseWriter, req *http.Request) {

	data := PageData{
		Title:   "Welcome to my website",
		Message: "Hello, World!",
	}

	t, _ := template.ParseFiles("HomePage.html")
	t.Execute(res, data)
}

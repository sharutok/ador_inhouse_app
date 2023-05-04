package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) Router() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.AppLink)
	router.HandlerFunc(http.MethodGet, "/index/of/ador/in-house/app", app.linkage)

	return router
}

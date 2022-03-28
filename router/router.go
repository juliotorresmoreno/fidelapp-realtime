package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/fidelapp-realtime/controllers"
	"github.com/juliotorresmoreno/fidelapp-realtime/middlewares"
)

func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middlewares.LoggerMiddleware)
	router.PathPrefix("/").Handler(controllers.NewHomeController())

	return router
}

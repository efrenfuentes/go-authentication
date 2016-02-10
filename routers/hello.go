package routers

import (
	"github.com/efrenfuentes/go-authentication/controllers"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/v1/hello", controllers.HelloIndex).Methods("GET")
	router.HandleFunc("/api/v1/hello/{name}", controllers.HelloName).Methods("GET")

	return router
}
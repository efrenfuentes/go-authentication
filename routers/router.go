package routers

import (
	"github.com/efrenfuentes/go-authentication/core"
	"github.com/gorilla/mux"
	"net/http"
)

func SetRoute(router *mux.Router, path, name, method string, handlerFunc func(w http.ResponseWriter, r *http.Request) ) *mux.Router {
	var handler http.HandlerFunc

	handler = handlerFunc
	handler = core.Logger(handler, name)
	router.HandleFunc(path, handler).Methods(method)

	return router
}

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	return router
}

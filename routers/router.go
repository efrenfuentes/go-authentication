package routers

import (
	"github.com/efrenfuentes/go-authentication/core/logger"
	"github.com/gorilla/mux"
	"net/http"
)

func SetRoute(router *mux.Router, path, name, method string, handlerFunc func(w http.ResponseWriter, r *http.Request) ) *mux.Router {
	var handler http.HandlerFunc

	handler = handlerFunc
	handler = logger.Logger(handler, name)
	router.HandleFunc(path, handler).Methods(method)

	return router
}

func Init() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	return router
}

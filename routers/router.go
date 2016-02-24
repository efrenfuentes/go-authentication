package routers

import (
	"net/http"

	"github.com/efrenfuentes/go-authentication/core/logger"
	"github.com/gorilla/mux"
)

func SetRoute(router *mux.Router, path, name, method string, handlerFunc func(w http.ResponseWriter, r *http.Request)) *mux.Router {
	var handler http.HandlerFunc

	handler = handlerFunc
	handler = logger.Logger(handler, name)
	router.HandleFunc(path, handler).Methods(method)

	return router
}

func Init() *mux.Router {
	router := mux.NewRouter()
	router = SetUsersRoutes(router)
	return router
}

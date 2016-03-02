package routers

import (
	"github.com/efrenfuentes/go-authentication/controllers"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {

	SetRoute(router, "/api/v1/hello", "HelloIndex", "GET", controllers.HelloIndex)
	SetProtectedRoute(router, "/api/v1/hello/{name}", "HelloName", "GET", controllers.HelloName)

	return router
}

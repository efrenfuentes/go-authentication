package routers

import (
	"github.com/efrenfuentes/go-authentication/controllers"
	"github.com/gorilla/mux"
)

func SetClientsRoutes(router *mux.Router) *mux.Router {

	SetProtectedRoute(router, "/api/v1/clients", "GetAllClients", "GET", controllers.GetAllClients)
	SetProtectedRoute(router, "/api/v1/clients/{id}", "GetClient", "GET", controllers.GetClient)
	SetProtectedRoute(router, "/api/v1/clients", "CreateClient", "POST", controllers.CreateClient)
	SetProtectedRoute(router, "/api/v1/clients/{id}", "UpdateClient", "PUT", controllers.UpdateClient)
	SetProtectedRoute(router, "/api/v1/clients/{id}", "DeleteClient", "DELETE", controllers.DeleteClient)

	return router
}

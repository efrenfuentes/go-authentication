package routers

import (
	"github.com/efrenfuentes/go-authentication/controllers"
	"github.com/gorilla/mux"
)

func SetGroupsRoutes(router *mux.Router) *mux.Router {

	SetProtectedRoute(router, "/api/v1/groups", "GetAllGroups", "GET", controllers.GetAllGroups)
	SetProtectedRoute(router, "/api/v1/groups/{id}", "GetGroup", "GET", controllers.GetGroup)
	SetProtectedRoute(router, "/api/v1/groups", "CreateGroup", "POST", controllers.CreateGroup)
	SetProtectedRoute(router, "/api/v1/groups/{id}", "UpdateGroup", "PUT", controllers.UpdateGroup)
	SetProtectedRoute(router, "/api/v1/groups/{id}", "DeleteGroup", "DELETE", controllers.DeleteGroup)

	return router
}

package routers

import (
	"github.com/efrenfuentes/go-authentication/controllers"
	"github.com/gorilla/mux"
)

func SetAbilitiesRoutes(router *mux.Router) *mux.Router {

	SetProtectedRoute(router, "/api/v1/abilities", "GetAllAbilities", "GET", controllers.GetAllAbilities)
	SetProtectedRoute(router, "/api/v1/abilities/{id}", "GetAbility", "GET", controllers.GetAbility)
	SetProtectedRoute(router, "/api/v1/abilities", "CreateAbility", "POST", controllers.CreateAbility)
	SetProtectedRoute(router, "/api/v1/abilities/{id}", "UpdateAbility", "PUT", controllers.UpdateAbility)
	SetProtectedRoute(router, "/api/v1/abilities/{id}", "DeleteAbility", "DELETE", controllers.DeleteAbility)

	return router
}

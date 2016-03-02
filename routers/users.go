package routers

import (
	"github.com/efrenfuentes/go-authentication/controllers"
	"github.com/gorilla/mux"
)

func SetUsersRoutes(router *mux.Router) *mux.Router {

	SetProtectedRoute(router, "/api/v1/users", "GetAllUsers", "GET", controllers.GetAllUsers)
	SetProtectedRoute(router, "/api/v1/users/{id}", "GetUser", "GET", controllers.GetUser)
	SetProtectedRoute(router, "/api/v1/users", "CreateUser", "POST", controllers.CreateUser)
	SetProtectedRoute(router, "/api/v1/users/{id}", "UpdateUser", "PUT", controllers.UpdateUser)
	SetProtectedRoute(router, "/api/v1/users/{id}", "DeleteUser", "DELETE", controllers.DeleteUser)

	SetRoute(router, "/api/v1/users/authenticate", "AuthenticateUser", "POST", controllers.AuthenticateUser)

	return router
}

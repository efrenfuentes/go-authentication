package routers

import (
	"github.com/efrenfuentes/go-authentication/controllers"
	"github.com/gorilla/mux"
)

func SetUsersRoutes(router *mux.Router) *mux.Router {

	SetRoute(router, "/api/v1/users", "GetAllUsers", "GET", controllers.GetAllUsers)
	SetRoute(router, "/api/v1/users/{id}", "GetUser", "GET", controllers.GetUser)
	SetRoute(router, "/api/v1/users", "CreateUser", "POST", controllers.CreateUser)
	SetRoute(router, "/api/v1/users/{id}", "UpdateUser", "PUT", controllers.UpdateUser)
	SetRoute(router, "/api/v1/users/{id}", "DeleteUser", "DELETE", controllers.DeleteUser)

	SetRoute(router, "/api/v1/users/authenticate", "AuthenticateUser", "POST", controllers.AuthenticateUser)

	return router
}

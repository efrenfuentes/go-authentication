package core

import (
	"github.com/efrenfuentes/go-authentication/routers"
	"github.com/efrenfuentes/go-authentication/core/settings"
	"net/http"
	"fmt"
)

type Server struct {}

func (s *Server) Run() {
	settings.Init()
	routes := routers.InitRoutes()

	mySettings := settings.Get()
	ip := mySettings["server"].(map[string]interface{})["ip"].(string)
	port := mySettings["server"].(map[string]interface{})["port"].(string)

	listen_in := ip + ":" + port

	fmt.Printf("Starting server in %s [%s]\n",
		listen_in,
		settings.GetEnvironment())

	http.Handle("/", routes)
	http.ListenAndServe(listen_in, nil)
}

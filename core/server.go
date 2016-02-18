package core

import (
	"fmt"
	"net/http"

	"github.com/efrenfuentes/go-authentication/core/settings"
	"github.com/efrenfuentes/go-authentication/database"
	"github.com/efrenfuentes/go-authentication/routers"
)

type Server struct{}

func (s *Server) Run() {
	settings.Init()
	database.Init(settings.Get()["database"].(map[string]interface{}))
	routes := routers.Init()

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

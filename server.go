package main

import (
	"github.com/efrenfuentes/go-authentication/routers"
	"github.com/efrenfuentes/go-utils/settings"
	"net/http"
	"fmt"
)

func main() {
	settings.Init()
	routes := routers.InitRoutes()

	mySettings := settings.Get()
	ip := mySettings["server"].(map[string]interface{})["ip"].(string)
	port := mySettings["server"].(map[string]interface{})["port"].(string)

	listen_in := ip + ":" + port

	fmt.Println("Starting server in " + listen_in)
	http.Handle("/", routes)

	http.ListenAndServe(listen_in, nil)
}

package main

import (
	"flag"

	"github.com/efrenfuentes/go-authentication/core"
	"github.com/efrenfuentes/go-authentication/core/keypair"
)

var keyFile *string
var migration *bool
var server *bool

func init() {
	keyFile = flag.String("keys", "", "generate rsa keys")
	migration = flag.Bool("migration", false, "creates data tables and records for admin user, base client, groups and abilities")
	server = flag.Bool("server", false, "start go-authenticate server")
}

func main() {
	flag.Parse()

	if *keyFile != "" {
		keypair.GenerateKeys(*keyFile, 2048)
	} else if *migration {
		var server core.Server

		server.Migrations()
	} else if *server {
		var server core.Server

		server.Run()
	}
}

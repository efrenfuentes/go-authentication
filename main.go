package main

import (
	"flag"

	"github.com/efrenfuentes/go-authentication/core"
	"github.com/efrenfuentes/go-authentication/core/keypair"
)

var keyFile *string

func init() {
	keyFile = flag.String("k", "", "generate rsa keys")
}

func main() {
	flag.Parse()

	if *keyFile == "" {
		var server core.Server

		server.Run()
	} else {
		keypair.GenerateKeys(*keyFile, 2048)
	}
}

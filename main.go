package main

import (
	"github.com/munjeli/testAPI/api"
	"log"
	"net/http"
)

func main() {

	certPath := "/etc/count/tls/server.crt"
	keyPath := "/etc/count/tls/server.key"
	http.HandleFunc("/count", api.CreateCount)
	log.Fatal(http.ListenAndServeTLS(":443", certPath, keyPath, nil))
}

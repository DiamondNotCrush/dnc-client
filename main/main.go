package main

import (
	"log"
	"net/http"

	"github.com/dnc/dnc-client/router"
)

func main() {
	log.Println("Listening on port " + router.Port())
	http.ListenAndServe(":"+router.Port(), router.Routes())
}

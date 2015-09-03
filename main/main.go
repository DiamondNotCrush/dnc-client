package main

import (
	"log"
	"net/http"

	"github.com/dnc/dnc-client/router"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	log.Println("Listening on port " + router.Port())
	open.Start("http://localhost:" + router.Port())
	http.ListenAndServe(":"+router.Port(), router.Routes())
}

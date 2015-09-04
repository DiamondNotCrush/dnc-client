package main

import (
	"log"
	"net/http"

	"github.com/dnc/dnc-client/helper"
	"github.com/dnc/dnc-client/router"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	helper.MakeConfig()
	log.Println("Listening on port " + router.Port())
	open.Start("http://127.0.0.1:" + router.Port())
	http.ListenAndServe(":"+router.Port(), router.Routes())
}

package main

import (
	"log"
	"net/http"

	"github.com/DiamondNotCrush/dnc-client/router"
	"github.com/DiamondNotCrush/dnc-client/router/info"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	port := info.Port()
	log.Println("Listening on port " + port)
	open.Start("http://127.0.0.1:" + port)
	http.ListenAndServe(":"+port, router.Routes())
}

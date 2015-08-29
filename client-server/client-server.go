package main

import (
  "log"
  "net/http"
  "github.com/dnc/dnc-client/helper"
)

func main() {
  http.HandleFunc("/", helper.ListOfFiles)
  http.HandleFunc("/mpd/", helper.FindMpd)

  log.Println("Listening on port 3000")
  http.ListenAndServe(":3000", nil)
}

package main

import (
  "log"
  "net/http"
  "github.com/dnc/dnc-client/helper"
)

func main() {
  http.HandleFunc("/", helper.ListOfFiles)
  // http.HandleFunc("/media/", func())
  fileServer := http.FileServer(http.Dir("../media"))
  http.Handle("/media/", http.StripPrefix("/media/", fileServer))
  
  log.Println("Listening on port 3000")
  http.ListenAndServe(":3000", nil)
}
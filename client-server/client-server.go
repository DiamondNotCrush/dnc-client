package main

import (
  "fmt"
  "log"
  "strings"
  "net/http"
  "encoding/json"
  "github.com/dnc/dnc-client/helper"
)

func main() {
  dir := "../mpd/"
  sharedFiles := helper.ListFiles(dir)

  // respond with shared files
  http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
    sharedFiles = helper.ListFiles(dir)
    js, err := json.Marshal(sharedFiles)
    helper.Check(err)
    res.Header().Set("Content-Type", "application/json")
    res.Write(js)
  })

  // serves up files if they are in the sharedFiles
  http.HandleFunc("/mpd/", func(res http.ResponseWriter, req *http.Request) {
    path := dir+strings.Join(strings.Split(req.URL.Path, "/")[2:], "")
    if sharedFiles[path] {
      fmt.Printf("Serving File: ")
      fmt.Println(path)
      http.ServeFile(res, req, path)
    } else {
      fmt.Printf("Blocking File: ")
      fmt.Println(path)
    }
  })

  log.Println("Listening on port 3000")
  http.ListenAndServe(":3000", nil)
}

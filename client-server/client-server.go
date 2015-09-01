package main

import (
  "log"
  "strings"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "github.com/dnc/dnc-client/helper"
)

func main() {
  // read port and dir from config file
  config, err := ioutil.ReadFile("./config")
  helper.Check(err)
  configArr := strings.Split(string(config), "\"")
  port := configArr[1]
  dir := configArr[3]
  if string(dir[len(dir) - 1]) != "/" {
    dir += "/"
  }

  // build initial library
  sharedFiles := helper.ListFiles(dir)

  // respond with shared files
  http.HandleFunc("/library/", func(res http.ResponseWriter, req *http.Request) {
    sharedFiles = helper.ListFiles(dir)
    js, err := json.Marshal(sharedFiles)
    helper.Check(err)
    res.Header().Set("Content-Type", "application/json")
    res.Header().Set("Access-Control-Allow-Origin", "*")
    res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    if req.Method == "OPTIONS" {
        return
    }
    log.Print("Sending library")
    res.Write(js)
  })

  // serves up files if they are in the sharedFiles
  http.HandleFunc("/shared/", func(res http.ResponseWriter, req *http.Request) {
    path := strings.Join(strings.Split(req.URL.Path, "/")[2:], "/")
    if sharedFiles[path] {
      log.Print("Serving file: "+path)
      res.Header().Set("Access-Control-Allow-Origin", "*")
      res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
      res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
      if req.Method == "OPTIONS" {
        return
      }
      http.ServeFile(res, req, dir+path)
    } else {
      log.Print("Blocking file: "+path)
    }
  })

  // start server
  log.Println("Listening on port "+port)
  http.ListenAndServe(":"+port, nil)
}

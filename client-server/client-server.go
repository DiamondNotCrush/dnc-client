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
  fmt.Print("Folder to share (eg: ./library): ")
  var input string
  fmt.Scanln(&input)
  if string(input[len(input) - 1]) != "/" {
    input += "/"
  }
  dir := input // ../directory/
  sharedFiles := helper.ListFiles(dir)

  // respond with shared files
  http.HandleFunc("/library/", func(res http.ResponseWriter, req *http.Request) {
    sharedFiles = helper.ListFiles(dir)
    js, err := json.Marshal(sharedFiles)
    helper.Check(err)
    res.Header().Set("Content-Type", "application/json")
    res.Write(js)
  })

  // serves up files if they are in the sharedFiles
  http.HandleFunc("/shared/", func(res http.ResponseWriter, req *http.Request) {
    path := strings.Join(strings.Split(req.URL.Path, "/")[2:], "/")
    if sharedFiles[path] {
      fmt.Print("Serving file: ")
      fmt.Println(path)
      http.ServeFile(res, req, dir+path)
    } else {
      fmt.Print("Blocking file: ")
      fmt.Println(path)
    }
  })

  log.Println("Listening on port 3000")
  http.ListenAndServe(":3000", nil)
}

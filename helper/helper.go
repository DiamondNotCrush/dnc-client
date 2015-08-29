package helper

import (
  "os"
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "github.com/zencoder/go-dash/mpd"
)

func check(err error) {
    if err != nil {
      panic(err)
    }
}

func ListOfFiles(res http.ResponseWriter, req *http.Request) {
  files, err := ioutil.ReadDir("../media")
  check(err)
  var fileArray []string
  for _, file := range files {
    fileArray = append(fileArray, file.Name())
  }
  js, err := json.Marshal(fileArray)
  res.Header().Set("Content-Type", "application/json")
  res.Write(js)
}

func Exists(name string) bool {
  if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
      return false
    }
  }
  return true
}

func FindMpd(res http.ResponseWriter, req *http.Request) {
  if Exists("../mpd"+req.URL.Path) {
    fmt.Println("Convert File: "+req.URL.Path)
  } else {
    fmt.Println(req.URL.Path)
    http.ServeFile(res, req, ".."+req.URL.Path)
  }
}

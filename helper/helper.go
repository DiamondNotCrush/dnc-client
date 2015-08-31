package helper

import (
  "os"
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"
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
  _, err := os.Stat(name)
  if err != nil {
    return false
  }
  return true
}

func FindMpd(res http.ResponseWriter, req *http.Request) {
  path := ".." + req.URL.Path
  if !Exists(path) {
    fmt.Printf("Convert File: ")
    fmt.Println(path)
  } else {
    fmt.Printf("Play File: ")
    fmt.Println(path)
    http.ServeFile(res, req, path)
  }
}

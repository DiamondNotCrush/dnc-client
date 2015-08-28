package helper

import (
  "fmt"
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
  for _, file := range files { 
    fmt.Println(file.Name())
  }
}

// func MediaFolder(res http.ResponseWriter, req *http:Request) {
//   http.ServeFile()
// }


    // fmt.Printf("Please input which directory\nwhat you want to share, start with \"/\":\n")
    // fmt.Scanf("%s",&dir)
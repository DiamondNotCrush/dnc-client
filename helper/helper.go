package helper

import (
  "io/ioutil"
)

func Check(err error) {
    if err != nil {
      panic(err)
    }
}

func listRecursion(dir string, fileObj map[string]bool) {
  files, err := ioutil.ReadDir(dir)
  Check(err)
  for _, file := range files {
    if file.IsDir() {
      listRecursion(dir+file.Name(), fileObj)
    } else {
      fileObj[dir+file.Name()] = true
    }
  }
}

func ListFiles(dir string) (map[string]bool) {
  var fileObj = make(map[string]bool)
  listRecursion(dir, fileObj)
  return fileObj
}

package helper

import (
  "io/ioutil"
)

func Check(err error) {
    if err != nil {
      panic(err)
    }
}

func listRecursion(dir string, localDir string, fileObj map[string]bool) {
  files, err := ioutil.ReadDir(dir+localDir)
  Check(err)
  for _, file := range files {
    if file.IsDir() {
      listRecursion(dir, localDir+file.Name()+"/", fileObj)
    } else {
      fileObj[localDir+file.Name()] = true
    }
  }
}

func ListFiles(dir string) (map[string]bool) {
  fileObj := make(map[string]bool)
  localDir := ""
  listRecursion(dir, localDir, fileObj)
  return fileObj
}

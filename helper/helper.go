package helper

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

//acceptable filetypes
var fileTypes = map[string]bool{
	"3gp":  true,
	"avi":  true,
	"mov":  true,
	"mp4":  true,
	"m4v":  true,
	"m4a":  true,
	"mp3":  true,
	"mkv":  true,
	"ogv":  true,
	"ogm":  true,
	"ogg":  true,
	"oga":  true,
	"webm": true,
	"wav":  true,
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

//lists all sub folders within the shared directory
func listRecursion(dir string, localDir string, fileObj map[string]bool) {
	files, err := ioutil.ReadDir(dir + localDir)
	Check(err)
	for _, file := range files {
		if file.IsDir() {
			listRecursion(dir, localDir+file.Name()+"/", fileObj)
		} else {
			strArr := strings.Split(file.Name(), ".")
			if fileTypes[strArr[len(strArr)-1]] {
				fileObj[localDir+file.Name()] = true
			}
		}
	}
}

//lists qualified files
func ListFiles(dir string) map[string]bool {
	fileObj := make(map[string]bool)
	localDir := ""
	listRecursion(dir, localDir, fileObj)
	return fileObj
}
//form data into JSON object for login/signup
func JSONify(str string) []byte {
	obj := make(map[string]string)
	strArr := strings.Split(str, "&")
	for i := range strArr {
		tuple := strings.Split(strArr[i], "=")
		obj[tuple[0]] = tuple[1]
	}
	js, err := json.Marshal(obj)
	Check(err)
	return js
}

func CheckAddr(addr string) bool {
	if strings.Split(addr, ":")[0] == "127.0.0.1" {
		return true
	} else {
		return false
	}
}

func MakeConfig() {
	if _, err := os.Stat("config"); err != nil {
		err := ioutil.WriteFile("config", []byte("dir=./&port=3000"), 0777)
		Check(err)
	}
}

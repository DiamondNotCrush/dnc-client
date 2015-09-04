package share

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dnc/dnc-client/router/info"
	"github.com/gorilla/mux"
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
var sharedFiles = ListFiles(info.Dir())

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
	listRecursion(info.Dir(), localDir, fileObj)
	return fileObj
}

func Library(res http.ResponseWriter, req *http.Request) {
	sharedFiles = ListFiles(info.Dir())
	js, err := json.Marshal(sharedFiles)
	Check(err)
	res.Header().Set("Content-Type", "application/json")
	res = info.SetCORS(res)
	if req.Method == "OPTIONS" {
		return
	}
	log.Print("Sending library")
	res.Write(js)
}

func Shared(res http.ResponseWriter, req *http.Request) {
	path := mux.Vars(req)["path"]
	res = info.SetCORS(res)
	if req.Method == "OPTIONS" {
		return
	}
	if sharedFiles[path] {
		log.Print("Serving file: " + path)
		http.ServeFile(res, req, info.Dir()+path)
	} else {
		log.Print("Blocking file: " + path)
		res.WriteHeader(404)
	}
}

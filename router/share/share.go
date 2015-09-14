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
var fileTypes = map[string]string{
	"3gp":  "v",
	"avi":  "v",
	"mov":  "v",
	"mp4":  "v",
	"m4v":  "v",
	"m4a":  "a",
	"mp3":  "a",
	"mkv":  "v",
	"ogv":  "v",
	"ogm":  "v",
	"ogg":  "v",
	"oga":  "a",
	"webm": "v",
	"wav":  "a",
}
var sharedFiles = ListFiles(info.Dir())

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getImdb(name string) string {
	res, err := http.Get("http://www.omdbapi.com/?t=" + name)
	check(err)
	body, err := ioutil.ReadAll(res.Body)
	filmData := make(map[string]string)
	err = json.Unmarshal(body, &filmData)
	check(err)
	poster, ok := filmData["Poster"]
	if ok {
		return poster
	} else {
		return ""
	}
}

//lists all sub folders within the shared directory
func listRecursion(dir string, localDir string, fileObj map[string]string) {
	files, err := ioutil.ReadDir(dir + localDir)
	check(err)
	for _, file := range files {
		if file.IsDir() {
			listRecursion(dir, localDir+file.Name()+"/", fileObj)
		} else {
			fileNameArr := strings.Split(file.Name(), ".")
			format := fileTypes[fileNameArr[len(fileNameArr)-1]]
			name := strings.Join(fileNameArr[:(len(fileNameArr)-1)], ".")
			if format == "v" {
				fileObj[localDir+file.Name()] = getImdb(name)
			} else if format == "a" {
				fileObj[localDir+file.Name()] = ""
			}
		}
	}
}

//lists qualified files
func ListFiles(dir string) map[string]string {
	fileObj := make(map[string]string)
	localDir := ""
	listRecursion(info.Dir(), localDir, fileObj)
	return fileObj
}

func Library(res http.ResponseWriter, req *http.Request) {
	// sharedFiles = ListFiles(info.Dir())
	js, err := json.Marshal(sharedFiles)
	check(err)
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
	if _, ok := sharedFiles[path]; ok {
		log.Print("Serving file: " + path)
		http.ServeFile(res, req, info.Dir()+path)
	} else {
		log.Print("Blocking file: " + path)
		res.WriteHeader(404)
	}
}

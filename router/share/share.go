package share

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/codeskyblue/go-sh"
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

var streaming = false

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//lists all sub folders within the shared directory
func listRecursion(dir string, localDir string, fileObj map[string]bool) {
	files, err := ioutil.ReadDir(dir + localDir)
	check(err)
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
	if sharedFiles[path] {
		if streaming {
			_ = sh.Command("taskkill", "/im", "vlc.exe").Run()
		}

		log.Print("Serving file: " + path)

		go func() {
			//Start ffmpeg output to new temporary
			streaming = true
			err := sh.Command("vlc", "-vvv", info.Dir()+path, "--sout", "#transcode{vcodec=VP80,vb=2000,acodec=vorb,ab=128,channels=2,samplerate=44100}:http{mux=webm,dst=:5000}").Run()

			if err != nil {
				log.Printf("Error transcoding: %s", err)
			}
		}()

		//Sleep so file is ready
		time.Sleep(1000 * 1000 * 1000 * 2)

		http.Redirect(res, req, "/stream", 301)
		return
	} else {
		log.Print("Blocking file: " + path)
		res.WriteHeader(404)
	}
}

package share

import (
	"encoding/json"
	"io/ioutil"
	"bufio"
	"os"
	"log"
	"time"
	"net/http"
	"strings"

	"github.com/dnc/dnc-client/router/info"
	"github.com/gorilla/mux"

	"github.com/codeskyblue/go-sh"	
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
		  /*
		      FFMpeg options
		        -i: input filename
		        -strict -2: force use of exprimental aac decoder
		        -movflags faststart: moves metadata from end of file to the beginning to decode and view at destination
		  */
		go func() {
			//Start ffmpeg output to new temporary
			err := sh.Command("ffmpeg", "-i", info.Dir()+path, "-strict", "-2", "-movflags", "faststart", "-preset", "ultrafast", info.Dir()+"temp_output.mp4").Run()		 
			//Check for errors in transcoding
			if err != nil {
			 				log.Printf("Error transcoding: %s", err)
			}
		}()

		//Sleep so file is ready
		time.Sleep(1000 * 1000 * 1000 * 5)

		hj, ok := res.(http.Hijacker)
		if !ok {
			http.Error(res, "Server does not support hijacking", http.StatusInternalServerError)
			return
		}

		conn, bufrw, err := hj.Hijack()
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		//Read from file
		file, err := os.Open(info.Dir()+"temp_output.mp4")
		if err != nil {
			log.Printf("Error reading file: %s", err)
		}

		b := bufio.NewReader(file)

		defer conn.Close()
		
		//stream to browser
		res.Header().Set("Content-Type", "video/mp4")

		if _, err := b.WriteTo(res); err != nil {
			log.Printf("Error writing to http response: %s", err)
		}
		
		bufrw.Flush()

		res.Write([]byte("File Sent"))

		//Remove file after it's been served
		os.Remove(info.Dir()+"temp_output.mp4")

	} else {
		log.Print("Blocking file: " + path)
		res.WriteHeader(404)
	}
}
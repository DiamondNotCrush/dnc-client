package router

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dnc/dnc-client/helper"
	"github.com/dnc/dnc-client/portal"
	"github.com/gorilla/mux"
)

var dir = getDir()
var sharedFiles = helper.ListFiles(dir)
var verify = false
var userid = -1

func setCORS(res http.ResponseWriter) http.ResponseWriter {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	return res
}

func getDir() string {
	config, err := ioutil.ReadFile("./config")
	helper.Check(err)
	configArr := strings.Split(string(config), "\"")
	dir := configArr[3]
	if string(dir[len(dir)-1]) != "/" {
		dir += "/"
	}
	return dir
}

func Port() string {
	config, err := ioutil.ReadFile("./config")
	helper.Check(err)
	configArr := strings.Split(string(config), "\"")
	return configArr[1]
}

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if !helper.CheckAddr(req.RemoteAddr) {
			return
		}
		portal.MainPage(res, req, dir, Port(), verify, userid)
	}).Methods("GET")

	// router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	if !helper.CheckAddr(req.RemoteAddr) {
	// 		return
	// 	}
	// 	data, err := ioutil.ReadAll(req.Body)
	// 	helper.Check(err)

	// 	portal.MainPage(res, req, dir, Port(), verify, userid)
	// }).Methods("POST")

	router.HandleFunc("/signup", func(res http.ResponseWriter, req *http.Request) {
		if !helper.CheckAddr(req.RemoteAddr) {
			return
		}
		portal.Signup(res, req)
	}).Methods("GET")

	router.HandleFunc("/signup", func(res http.ResponseWriter, req *http.Request) {
		if !helper.CheckAddr(req.RemoteAddr) {
			return
		}
		data, err := ioutil.ReadAll(req.Body)
		helper.Check(err)
		js := bytes.NewReader(helper.JSONify(string(data)))
		sres, err := http.Post("http://dnctest.herokuapp.com/user/addUser", "application/json", js)
		helper.Check(err)
		sdata, err := ioutil.ReadAll(sres.Body)
		helper.Check(err)
		var obj map[string]*json.RawMessage
		err = json.Unmarshal(sdata, &obj)
		helper.Check(err)
		err = json.Unmarshal(*obj["id"], &userid)
		helper.Check(err)
		if err != nil {
			log.Println("Signup failed")
		} else {
			log.Println("Signup success")
		}
		http.Redirect(res, req, "/", 302)
	}).Methods("POST")

	router.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
		if !helper.CheckAddr(req.RemoteAddr) {
			return
		}
		portal.Login(res, req)
	}).Methods("GET")

	router.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
		if !helper.CheckAddr(req.RemoteAddr) {
			return
		}
		data, err := ioutil.ReadAll(req.Body)
		helper.Check(err)
		js := bytes.NewReader(helper.JSONify(string(data)))
		sres, err := http.Post("http://dnctest.herokuapp.com/user/login", "application/json", js)
		helper.Check(err)
		sdata, err := ioutil.ReadAll(sres.Body)
		helper.Check(err)
		var obj map[string]*json.RawMessage
		err = json.Unmarshal(sdata, &obj)
		helper.Check(err)
		err = json.Unmarshal(*obj["id"], &userid)
		helper.Check(err)
		if err != nil {
			log.Println("Login failed")
		} else {
			log.Println("Login success")
		}
		http.Redirect(res, req, "/", 302)
	}).Methods("POST")

	router.HandleFunc("/verify", func(res http.ResponseWriter, req *http.Request) {
		verify = true
		res.WriteHeader(200)
	}).Methods("GET")

	router.HandleFunc("/library", func(res http.ResponseWriter, req *http.Request) {
		sharedFiles = helper.ListFiles(dir)
		js, err := json.Marshal(sharedFiles)
		helper.Check(err)
		res.Header().Set("Content-Type", "application/json")
		res = setCORS(res)
		if req.Method == "OPTIONS" {
			return
		}
		log.Print("Sending library")
		res.Write(js)
	}).Methods("GET")

	router.HandleFunc("/shared/{path:.*}", func(res http.ResponseWriter, req *http.Request) {
		path := mux.Vars(req)["path"]
		res = setCORS(res)
		if req.Method == "OPTIONS" {
			return
		}
		if sharedFiles[path] {
			log.Print("Serving file: " + path)
			http.ServeFile(res, req, dir+path)
		} else {
			log.Print("Blocking file: " + path)
			res.WriteHeader(404)
		}
	}).Methods("GET")

	return router
}

package info

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var dir = getDir()
var port = getPort()

//get shared folder from config file
func getDir() string {
	makeConfig()
	config, err := ioutil.ReadFile("config")
	check(err)
	dir := strings.Split(strings.Split(string(config), "&")[0], "=")[1]
	if string(dir[len(dir)-1]) != "/" {
		dir += "/"
	}
	return dir
}

//get port from config file
func getPort() string {
	makeConfig()
	config, err := ioutil.ReadFile("config")
	check(err)
	port := strings.Split(strings.Split(string(config), "&")[1], "=")[1]
	return port
}

func Dir() string {
	return dir
}

func Port() string {
	return port
}

func makeConfig() {
	if _, err := os.Stat("config"); err != nil {
		err := ioutil.WriteFile("config", []byte("dir=./&port=3000"), 0777)
		check(err)
	}
}

func SetCORS(res http.ResponseWriter) http.ResponseWriter {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	return res
}

func CheckAddr(addr string) bool {
	if strings.Split(addr, ":")[0] == "127.0.0.1" {
		return true
	} else {
		return false
	}
}

func ChangeConfig(res http.ResponseWriter, req *http.Request) {
	if !CheckAddr(req.RemoteAddr) {
		return
	}
	data, err := ioutil.ReadAll(req.Body)
	check(err)
	data = []byte(strings.Replace(string(data), "%2F", "/", -1))
	err = ioutil.WriteFile("config", data, 0777)
	check(err)
	dir = getDir()
	port = getPort()
	http.Redirect(res, req, "/", 302)
}

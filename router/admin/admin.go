package admin

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/DiamondNotCrush/dnc-client/portal"
	"github.com/DiamondNotCrush/dnc-client/router/info"
)

var verify = false
var userid = -1

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//form data into JSON object for login/signup
func jSONify(str string) []byte {
	obj := make(map[string]string)
	strArr := strings.Split(str, "&")
	for i := range strArr {
		tuple := strings.Split(strArr[i], "=")
		obj[tuple[0]] = tuple[1]
	}
	js, err := json.Marshal(obj)
	check(err)
	return js
}

func MainPage(res http.ResponseWriter, req *http.Request) {
	if !info.CheckAddr(req.RemoteAddr) {
		return
	}
	portal.MainPage(res, req, info.Dir(), info.Port(), verify, userid)
}

func GetSignup(res http.ResponseWriter, req *http.Request) {
	if !info.CheckAddr(req.RemoteAddr) {
		return
	}
	portal.Signup(res, req)
}

func PostSignup(res http.ResponseWriter, req *http.Request) {
	if !info.CheckAddr(req.RemoteAddr) {
		return
	}
	data, err := ioutil.ReadAll(req.Body)
	check(err)
	js := bytes.NewReader(jSONify(string(data)))
	sres, err := http.Post("http://dnctest.herokuapp.com/user/addUser", "application/json", js)
	check(err)
	if err != nil {
		log.Println("Signup failed")
	} else {
		log.Println("Signup success")
	}
	Ping(res, req, sres)
}

func GetLogin(res http.ResponseWriter, req *http.Request) {
	if !info.CheckAddr(req.RemoteAddr) {
		return
	}
	portal.Login(res, req)
}

func PostLogin(res http.ResponseWriter, req *http.Request) {
	if !info.CheckAddr(req.RemoteAddr) {
		return
	}
	data, err := ioutil.ReadAll(req.Body)
	check(err)
	js := bytes.NewReader(jSONify(string(data)))
	sres, err := http.Post("http://dnctest.herokuapp.com/user/login", "application/json", js)
	check(err)
	if err != nil {
		log.Println("Login failed")
	} else {
		log.Println("Login success")
	}
	Ping(res, req, sres)
}

func Ping(res http.ResponseWriter, req *http.Request, sres *http.Response) {
	sdata, err := ioutil.ReadAll(sres.Body)
	check(err)
	var obj map[string]*json.RawMessage
	err = json.Unmarshal(sdata, &obj)
	check(err)
	err = json.Unmarshal(*obj["id"], &userid)
	check(err)
	port, err := strconv.Atoi(info.Port())
	check(err)
	ping := map[string]int{
		"userid": userid,
		"port":   port,
	}
	jsstr, err := json.Marshal(ping)
	check(err)
	_, err = http.Post("http://dnctest.herokuapp.com/connection/addConnection", "application/json", bytes.NewReader(jsstr))
	check(err)
	http.Redirect(res, req, "/", 302)
}

func Verify(res http.ResponseWriter, req *http.Request) {
	verify = true
	res.WriteHeader(200)
}

package portal

import (
	"net/http"
	"text/template"
)

type Page struct {
	Dir  string
	Port string
}

var tmain, _ = template.ParseFiles("../src/github.com/dnc/dnc-client/portal/templates/index.html")
var tsignup, _ = template.ParseFiles("../src/github.com/dnc/dnc-client/portal/templates/signup.html")
var tlogin, _ = template.ParseFiles("../src/github.com/dnc/dnc-client/portal/templates/login.html")

func MainPage(res http.ResponseWriter, req *http.Request, dir string, port string) {
	page := Page{dir, port}
	tmain.Execute(res, page)
}

func Signup(res http.ResponseWriter, req *http.Request) {
	tsignup.Execute(res, nil)
}

func Login(res http.ResponseWriter, req *http.Request) {
	tlogin.Execute(res, nil)
}

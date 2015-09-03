package portal

import (
	"net/http"
	"text/template"
)

type Page struct {
	Dir         string
	Port        string
	LoginStatus string
	Verify      string
}

var tmain, _ = template.ParseFiles("../src/github.com/dnc/dnc-client/portal/templates/index.html")
var tsignup, _ = template.ParseFiles("../src/github.com/dnc/dnc-client/portal/templates/signup.html")
var tlogin, _ = template.ParseFiles("../src/github.com/dnc/dnc-client/portal/templates/login.html")

func MainPage(res http.ResponseWriter, req *http.Request, dir string, port string, result bool, userid int) {
	verify := `<p style="color:red">Unverified</p>`
	if result {
		verify = `<p style="color:green">Verification complete!</p>`
	}
	status := `<p style="color:red">Not logged in</p>`
	if userid > -1 {
		status = `<p style="color:green">Logged in!</p>`
	}
	page := Page{dir, port, status, verify}
	tmain.Execute(res, page)
}

func Signup(res http.ResponseWriter, req *http.Request) {
	tsignup.Execute(res, nil)
}

func Login(res http.ResponseWriter, req *http.Request) {
	tlogin.Execute(res, nil)
}

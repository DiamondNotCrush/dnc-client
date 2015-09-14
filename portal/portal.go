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
var tstream, _ = template.ParseFiles("../src/github.com/dnc/dnc-client/portal/templates/stream.html")

//check user status
func MainPage(res http.ResponseWriter, req *http.Request, dir string, port string, result bool, userid int) {
	verify := `<span style="color:red">Unverified</span>`
	if result {
		verify = `<span style="color:green">Verification complete!</span>`
	}
	status := `<span style="color:red">Not logged in</span>`
	if userid > -1 {
		status = `<span style="color:green">Logged in!</span>`
	}
	page := Page{dir, port, status, verify}
	tmain.Execute(res, page)
}

//redirect to signup html
func Signup(res http.ResponseWriter, req *http.Request) {
	tsignup.Execute(res, nil)
}

//redirect to login html
func Login(res http.ResponseWriter, req *http.Request) {
	tlogin.Execute(res, nil)
}

func Stream(res http.ResponseWriter, req *http.Request) {
	tstream.Execute(res, nil)
}

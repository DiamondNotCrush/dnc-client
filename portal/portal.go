package portal

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title       string
	Dir         string
	Port        string
	LoginStatus string
	Verify      string
}

var templates = template.Must(template.ParseFiles(
	"../src/github.com/dnc/dnc-client/portal/templates/header.html",
	"../src/github.com/dnc/dnc-client/portal/templates/footer.html",
	"../src/github.com/dnc/dnc-client/portal/templates/main.html",
	"../src/github.com/dnc/dnc-client/portal/templates/signup.html",
	"../src/github.com/dnc/dnc-client/portal/templates/login.html"))

// var tmain, _ = template.ParseFiles("../src/github.com/dnc/dnc-client/portal/templates/main.html")
// var tsignup, _ = template.ParseFiles("../src/github.com/dnc/dnc-client/portal/templates/signup.html")
// var tlogin, _ = template.ParseFiles("../src/github.com/dnc/dnc-client/portal/templates/login.html")

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
	// page := Page{dir, port, status, verify}
	templates.ExecuteTemplate(res, "main", &Page{"Home", dir, port, status, verify})
}

//redirect to signup html
func Signup(res http.ResponseWriter, req *http.Request) {
	data := struct{ Title string }{"Signup"}
	templates.ExecuteTemplate(res, "signup", data)
}

//redirect to login html
func Login(res http.ResponseWriter, req *http.Request) {
	data := struct{ Title string }{"Login"}
	templates.ExecuteTemplate(res, "login", data)
}

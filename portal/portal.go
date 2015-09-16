package portal

import (
	"net/http"

	"github.com/dnc/dnc-client/portal/portalTemplates"
)

var templates = portalTemplates.Generate()

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
	data := struct {
		Title       string
		Dir         string
		Port        string
		LoginStatus string
		Verify      string
	}{"Home", dir, port, status, verify}
	templates.ExecuteTemplate(res, "main", data)
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

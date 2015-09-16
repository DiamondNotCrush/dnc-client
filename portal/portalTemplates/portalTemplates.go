package portalTemplates

import "text/template"

const header string = `{{define "header"}}
<!DOCTYPE html>
<html>
  <head>
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
  </head>
  <body>
      <nav class="navbar navbar-default" role="navigation">
        <div class="navbar-header">
          <a class="navbar-brand" href="/">Main</a>
        </div>
        <div class="collapse navbar-collapse navbar-ex1-collapse">  
          <ul class="nav navbar-nav">
            <li><a href="/login">Login</a></li>
            <li><a href="/signup">Signup</a></li>
          </ul>
        </div>
      </nav>
    <div class="col-xs-12 col-sm-6 col-sm-offset-3 col-md-4 col-md-offset-4 ng-scope">
{{end}}`

const footer string = `{{define "footer"}}
    </div>
    <p class="navbar-text navbar-fixed-bottom">Diamond not Crush</p>  
  </body>
</html>
{{end}}`

const main string = `{{define "main"}}
{{template "header" .}}
    <h1>Admin Portal</h1>
    <h4>Status</h4>
    <dir class="well well-sm">
    <p>Login Status: {{.LoginStatus}}</p>
    <p>Connection Status: {{.Verify}}
    <br><br>
    <input type="button" onclick="location.href='/login';" value="Login" class="btn btn-default"/>
    <input type="button" onclick="location.href='/signup';" value="Signup" class="btn btn-default"/>
    </p></dir><br>
    <h4>Change Configuration</h4>
    <p>
    <form action="" method="post" class="well well-sm">
      Directory: {{.Dir}}<br>
      <input type="text" name="dir" class="form-control" value="{{.Dir}}">
      Port: {{.Port}}<br>
      <input type="text" name="port" class="form-control" value="{{.Port}}"><br>
      <input type="submit" value="Set" class="btn btn-default">
      </p><p>
      A port change requires a restart.
      <p>
    </form>
{{template "footer" .}}
{{end}}`

const signup string = `{{define "signup"}}
{{template "header" .}}
    <h4>Signup</h4>
    <form action="" method="post" class="well well-sm">
      <p>
      <input type="text" name="username" class="form-control" placeholder="Username"><br>
      <input type="email" name="email" class="form-control" placeholder="Email"><br>
      <input type="password" name="password" class="form-control" placeholder="Password"><br><br>
      <input type="button" onclick="location.href='/';" value="Back" class="btn btn-default"/>
      <input type="submit" value="Submit" class="btn btn-default">
      </p>
    </form>
{{template "footer" .}}
{{end}}`

const login string = `{{define "login"}}
{{template "header" .}}
    <h4>Login</h4>
    <p>
    <form action="" method="post" class="well well-sm">
      <input type="email" name="email" class="form-control" placeholder="Email"><br>
      <input type="password" name="password" class="form-control" placeholder="Password"><br><br>
      <input type="button" onclick="location.href='/';" value="Back" class="btn btn-default"/>
      <input type="submit" value="Submit" class="btn btn-default">
    </form>
    </p>
{{template "footer" .}}
{{end}}`

func Generate() *template.Template {
	collection := template.New("")
	_, err := collection.Parse(header)
	_, err = collection.Parse(footer)
	_, err = collection.Parse(main)
	_, err = collection.Parse(signup)
	_, err = collection.Parse(login)
	if err != nil {
		panic(err)
	}
	return collection
}

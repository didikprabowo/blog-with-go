package login

import (
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("templates/**/*.html"))

func Auth(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "login", tmpl)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// email := r.PostForm.Get("email")
	// fmt.Println(email)
	w.Write([]byte("hai"))
}

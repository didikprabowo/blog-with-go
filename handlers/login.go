package handlers

import (
	"fmt"
	"github.com/didikprabowo/blog/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("templates/**/*.html"))
var store = sessions.NewCookieStore([]byte("didikprabowo"))

func Auth(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "login", tmpl)
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	db := database.MySQL()
	var emailNews string
	var passwordNews string

	db.QueryRow("SELECT email, password FROM users WHERE email = ?", email).Scan(&emailNews, &passwordNews)
	var data = CheckPasswordHash(password, passwordNews)
	if data != true {
		w.Write([]byte("Email OR Password Invalid"))
		return
	}

	//save session

	sessionsLogin, _ := store.Get(r, "login")
	sessionsLogin.Values["email"] = email
	sessionsLogin.Values["login"] = true
	sessionsLogin.Save(r, w)
	fmt.Println(sessionsLogin.Values["login"])
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func Register(w http.ResponseWriter, r *http.Request) {
	email := "didikprab@gmail.com"
	var password string
	password, _ = HashPassword("didikprab@gmail.com")
	db := database.MySQL()
	_, err := db.Exec("INSERT INTO users(email, password) VALUES(?, ?)", email, password)
	if err != nil {
		w.Write([]byte("Invalid Insert"))
	}

}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//Logout
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login")
	session.Options.MaxAge = -1
	session.Save(r, w)
}

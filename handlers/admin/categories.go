package admin

import (
	"strings"
	// "fmt"
	"github.com/didikprabowo/blog/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"net/http"
	"text/template"
)

var store = sessions.NewCookieStore([]byte("didikprabowo"))
var tmpl = template.Must(template.ParseGlob("templates/**/*.html"))

type Meta struct {
	Title string
}
type Category struct {
	ID          int
	Name        string
	Description string
	Slug        string
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login")
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/auth", 301)

	}

	db := database.MySQL()
	categories, err := db.Query("SELECT id,name,description,slug FROM categories ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Category{}
	res := []Category{}
	for categories.Next() {
		var id int
		var name, description, slug string
		err := categories.Scan(&id, &name, &description, &slug)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.Name = strings.ToUpper(name)
		emp.Description = description
		emp.Slug = slug
		res = append(res, emp)
	}
	m := map[string]interface{}{
		"Results": res,
		"Titles":  Meta{Title: "Categories"},
	}
	tmpl.ExecuteTemplate(w, "category.html", m)
}

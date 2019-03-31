package admin

import (
	"github.com/didikprabowo/blog/database"
	"github.com/didikprabowo/blog/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"net/http"
	"strings"
	"text/template"
)

var store = sessions.NewCookieStore([]byte("didikprabowo"))
var tmpl = template.Must(template.ParseGlob("templates/**/*.html"))

type Meta struct {
	Title string
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
	emp := models.Category{}
	res := []models.Category{}
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
	defer db.Close()
	m := map[string]interface{}{
		"Results": res,
		"Titles":  Meta{Title: "Categories"},
	}
	tmpl.ExecuteTemplate(w, "category.html", m)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{
		"Titles": Meta{Title: "Create Categories"},
	}
	tmpl.ExecuteTemplate(w, "addcategory.html", m)
}
func StoreCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		slug := strings.ToLower(strings.Replace(name, " ", "-", -1))
		db := database.MySQL()

		query, err := db.Prepare("INSERT INTO categories (slug,name,description) values(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		query.Exec(slug, name, description)
		defer db.Close()
		http.Redirect(w, r, "/admin/category/", 301)
	} else {
		http.Redirect(w, r, "/admin/category/", 301)
	}
}

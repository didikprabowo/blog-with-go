package admin

import (
	"html/template"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	t := template.New("footer")
	t, _ = t.Parse("hello !")
	t.Execute(w, nil)
}

package admin

import (
	"html/template"
)

var tmpl = template.Must(template.ParseGlob("templates/**/*.html"))

// Meta
type Meta struct {
	Title string
}

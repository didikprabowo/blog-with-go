package web

import (
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("templates/**/*.html"))

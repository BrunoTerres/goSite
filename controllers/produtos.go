package controllers

import (
	"net/http"
	"std/bruno/goSite/models"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.SelectTodosProdutos()
	tmpl.ExecuteTemplate(w, "Index", produtos)
}

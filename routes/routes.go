package routes

import (
	"net/http"
	"std/bruno/goSite/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}

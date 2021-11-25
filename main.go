package main

import (
	"net/http"
	"std/bruno/goSite/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

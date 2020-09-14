package controller

import (
	"github.com/fikriaakbar/GoAPI/gocode/database"
	"net/http"

	"github.com/fikriaakbar/GoAPI/gocode/page"

	"github.com/gorilla/mux"
)

//ListController func
func ListController() {
	r := mux.NewRouter()
	r.HandleFunc("/", page.HomeHandler)
	r.HandleFunc("/db", database.SelectDbDate)
	r.HandleFunc("/indb", database.InsertDb)
	http.ListenAndServe(":9000", r)
}

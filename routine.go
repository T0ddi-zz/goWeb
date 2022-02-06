package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		vars["title"] // the book title slug
		vars["page"] // the page
	})

	http.ListenAndServe(":80", r)

}
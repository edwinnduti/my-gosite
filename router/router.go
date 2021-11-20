package router

import (
	"net/http"

	"github.com/edwinnduti/my-gosite/lib"
	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	//Register router
	r := mux.NewRouter()

	//handled routes
	r.HandleFunc("/", lib.HomeHandler).Methods("GET")
	r.HandleFunc("/forms/contact", lib.SendMail).Methods("POST")
	r.HandleFunc("/d2ip", lib.Domain2IP).Methods("GET")
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(lib.Dir))))

	return r
}

package main

import(
	"net/http"
	"os"
	"log"
	"html/template"
	"github.com/gorilla/mux"
)

var (
	templ = template.Must(template.ParseGlob("templates/*.html"))
)

func main(){
	//Register router
	r := mux.NewRouter()

	//handled routes
	r.HandleFunc("/",homeHandler).Methods("GET")
	r.HandleFunc("/folio",folioHandler).Methods("GET")

	//Get port 
//	Port := "8086"
	Port := os.Getenv("PORT")
	if Port == ""{
		Port == "8081"
	}

	//start server
	server := &http.Server{
		Handler: r,
		Addr   : ":"+Port,
	}

	//log output
	log.Printf("Listening on port: %v",Port)
	server.ListenAndServe()
}

//home handler
func homeHandler(w http.ResponseWriter,r *http.Request) {
	//fmt.Fprintf(w,"Welcome to home")
	err := templ.ExecuteTemplate(w,"index.html",nil)
	Check(err)
}

//folio handler
func folioHandler(w http.ResponseWriter,r *http.Request) {
	err := templ.ExecuteTemplate(w,"site.html",nil)
	Check(err)
}

//error handler
func Check(err error) {
	if err != nil{
		log.Fatalln(err)
	}
}

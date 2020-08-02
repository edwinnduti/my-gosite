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
	dir = "assets/"
)

func main(){
	//Register router
	r := mux.NewRouter()

	//handled routes
	r.HandleFunc("/",homeHandler).Methods("GET")
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(dir))))

	//Get port
	Port := os.Getenv("PORT")
	if Port == ""{
		Port = "8077"
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
	err := templ.ExecuteTemplate(w,"index.html",nil)
	Check(err)
}

//error handler
func Check(err error) {
	if err != nil{
		log.Fatalln(err)
	}
}

package main

import(
	"net/http"
	"net/smtp"
	"os"
	"log"
	"html/template"
	"github.com/gorilla/mux"
)

//credentials
var (
	TO_EMAIL = []string{"nduti316@gmail.com"}
	PASSWORD = "tkqhbdrjjxhkrkyr"
	USERNAME = "nduti316@gmail.com"
)

//smtp FQDN
type SmtpServer struct{
	Host string
	Port string
}

//message form
type Message struct{
	Name    string
	From    string
	Subject []byte
	Message []byte
}

//template folder
var (
	templ = template.Must(template.ParseGlob("templates/*.html"))
	dir = "assets/"
)

func main(){
	//Register router
	r := mux.NewRouter()

	//handled routes
	r.HandleFunc("/",homeHandler).Methods("GET")
	r.HandleFunc("/forms/contact",sendMail).Methods("POST")
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(dir))))

	//Get port
	Port := os.Getenv("PORT")
	if Port == ""{
		Port = "8090"
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

//mail handler
func sendMail(w http.ResponseWriter,r *http.Request) {
	message := &Message{
		Name  : r.PostFormValue("name"),
		From  : r.PostFormValue("email"),
		Subject : []byte(r.PostFormValue("subject")),
		Message : []byte(r.PostFormValue("message")),
	}


	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";

	subject := "Subject: Mail from your site!\n"
	name := "Message from : "+ message.Name
	from := "Email address: "+message.From

	if message.Subject == nil{
		message.Subject = []byte(subject)
	}

	msg := []byte(string(message.Subject)+mime+"<html><body><h1>"+string(message.Subject)+"</h1><br><h3>"+name+"</h3><br><h3>"+from+"</h3><br><pre>"+string(message.Message)+"</pre></body></html>")

	message.Message = msg

	smtpServer := &SmtpServer{
		Host : "smtp.gmail.com",
		Port : "587",
	}

	address := smtpServer.Host+":"+smtpServer.Port

	auth := smtp.PlainAuth("",USERNAME,PASSWORD,smtpServer.Host)

	err := smtp.SendMail(address,auth,message.From,TO_EMAIL,message.Message)
	if err!=nil{
		Check(err)
	}
	log.Println("Email Sent!")


}

//error handler
func Check(err error) {
	if err != nil{
		log.Fatalln(err)
	}
}

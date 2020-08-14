package main

import(
	"net/http"
	"net/smtp"
	"os"
	"log"
	"html/template"
	"github.com/gorilla/mux"
	"github.com/edwinnduti/devtoapi"
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

//Blogpost struct
type Blogposts struct{
	Items map[string]string
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
		Port = "8072"
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

	//GET title and descriptions
	titles,descriptions,err := devtoapi.GetTitles("nduti")
	Check(err)

	//map titles to descriptions
	items := make(map[string]string)
	for k,_ := range titles{
		items[titles[k]] = descriptions[k]
	}

	//map blogpost struct to items
	blogposts := &Blogposts{
		Items : items,
	}

	//render template
	err = templ.ExecuteTemplate(w,"index.html",blogposts)
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

	subject := "Subject :Mail from your site!\n"
	name := "Message from : "+ message.Name
	from := "Email address: "+message.From

	if message.Subject == nil{
		message.Subject = []byte(subject)
	}

	msg := []byte("From: "+string(message.From)+"\r\n"+"Subject:"+string(message.Subject)+"\r\n"+mime+"<html><head><style>#rcorners {border-radius: 25px; background: #8AC007; padding: 20px; width: 90%; height: 100%;}</style></head><body id=\"rcorners\"><h3 background-color=\"blue\">"+string(message.Subject)+"</h3><h3>"+name+"</h3><h3>"+from+"</h3><br><pre>"+string(message.Message)+"</pre></body></html>")

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

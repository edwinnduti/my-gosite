package model

//smtp FQDN
type SmtpServer struct {
	Host string
	Port string
}

//message form
type Message struct {
	Name    string
	From    string
	Subject []byte
	Message []byte
}

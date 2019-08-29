//definiciones de estrucutras para el envio de emails

package iemailstt

//AuthEmail econfiguración para loguearse en el envio de emails
type AuthEmail struct {
	Server   string
	Port     int
	User     string
	Password string
	SSL      bool
}

//EmailMessage Es el parametro que se le tiene que enviar al metodo SendMail
type EmailMessage struct {
	Subject  string
	Body     string
	From     string
	To       []string
	Cc       CC
	Attached AttachFile
}

//CC envia un email con copia
type CC struct {
	Email string
	Name  string
}

//AttachFile archivos adjuntos
type AttachFile struct {
	File string
}

//WrappForgot Utilizamos esta estructura para parsear los templantes
// y agregarle las variables correspondientes
type WrappForgot struct {
	Code     string
	Fullname string
}

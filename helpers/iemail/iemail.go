//clase para envio de emails
//https://godoc.org/gopkg.in/gomail.v2#example-package

package iemail

import (
	"crypto/tls"

	iemailstt "isystem/helpers/iemail/structs"

	"gopkg.in/gomail.v2"
)

//SendEamil Envía un email
func SendEamil(emailMessage iemailstt.EmailMessage) (bool, error) {
	conf := DefaultConfig()

	m := gomail.NewMessage()
	m.SetHeader("From", emailMessage.From)
	m.SetHeader("To", emailMessage.To...) //"bob@example.com", "cora@example.com"

	if emailMessage.Cc.Email != "" {
		m.SetAddressHeader("Cc", emailMessage.Cc.Email, emailMessage.Cc.Name) // "dan@example.com", "Dan"
	}

	m.SetHeader("Subject", emailMessage.Subject)
	m.SetBody("text/html", emailMessage.Body)

	if emailMessage.Attached.File != "" {
		m.Attach(emailMessage.Attached.File) //"/home/Alex/lolcat.jpg"
	}

	d := gomail.NewDialer(conf.Server, conf.Port, conf.User, conf.Password)
	d.SSL = conf.SSL
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false, ServerName: conf.Server}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return false, err
	}

	return true, nil
}

/*
Servidor de correo entrante (IMAP)
imap.gmail.com
Requiere SSL: Sí
Puerto: 993

Servidor de correo saliente (SMTP)
smtp.gmail.com
Requiere SSL: Sí
Requiere TLS: Sí (si está disponible)
Requiere autenticación: Sí
Puerto para SSL: 465
Puerto para TLS/STARTTLS: 587

Nombre completo o nombre visible	Tu nombre
Nombre de la cuenta, nombre de usuario o dirección de correo electrónico	Tu dirección de correo electrónico completa
Contraseña	Tu contraseña de Gmail
*/

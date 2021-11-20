package helper

import (
	"bytes"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
)

var auth smtp.Auth

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func NewRequest(to []string, subject string) *Request {
	return &Request{
		to:      to,
		subject: subject,
	}
}

func (r *Request) Send(templateName string, items interface{}) (bool, error) {
	err := r.ParseTemplate(templateName, items)
	if err != nil {
		log.Fatal(err)
	}
	if err := r.SendEmail(); err != nil {
		return false, err
	}

	return true, nil
}

func (r *Request) SendEmail() error {
	server := "mail.grupogit.com"
	port := ":26"
	email := "efrenherrera@grupogit.com"
	password := "EHerrera.1234"

	addr := server + port
	auth = smtp.PlainAuth("", email, password, server)

	from := mail.Address{"", email}
	to := mail.Address{"", r.to[0]}
	subj := "This is the email subject"
	body := "From: " + email + "\r\nTo: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	smtp.SendMail(addr, auth, email, r.to, []byte(body))

	// // TLS config
	// tlsconfig := &tls.Config{
	// 	InsecureSkipVerify: true,
	// 	ServerName:         server,
	// }
	// c, err := smtp.Dial(addr)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// c.StartTLS(tlsconfig)

	// // Auth
	// if err = c.Auth(auth); err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(auth)

	// // To && From
	// if err = c.Mail(from.Address); err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(from.Address)
	// if err = c.Rcpt(to.Address); err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(to.Address)
	// // Data
	// w, err := c.Data()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// _, err = w.Write([]byte(body))
	// if err != nil {
	// 	log.Panic(err)
	// }

	// err = w.Close()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// c.Quit()

	return nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

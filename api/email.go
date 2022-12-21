package handler

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	gomail "gopkg.in/mail.v2"
)

// Email send to email
func Email(w http.ResponseWriter, r *http.Request) {
	email := os.Getenv("GMAIL")
	pass := os.Getenv("PASS_GMAIL")
	if email == "" || pass == "" {
		log.Fatalln("configure sus variables de entorno")
		return
	}
	m := gomail.NewMessage()

	m.SetHeader("From", "luispfcanales@gmail.com")
	m.SetHeader("To", "lpfunoc@unamad.edu.pe")

	m.SetHeader("Subject", "Gophers GO!")

	t := template.Must(template.ParseFiles("WellcomeTemplate.html"))
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		return t.Execute(w, "Registrate")
	})

	d := gomail.NewDialer("smtp.gmail.com", 587, email, pass)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	w.Write([]byte("send email"))
}

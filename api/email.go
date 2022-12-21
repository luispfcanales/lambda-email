package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	gomail "gopkg.in/mail.v2"
)

type Person struct {
	Code  string `json:"code,omitempty"`
	Email string `json:"email,omitempty"`
}

// Email send to email
func Email(w http.ResponseWriter, r *http.Request) {
	p := &Person{}
	json.NewDecoder(r.Body).Decode(p)

	email := os.Getenv("GMAIL")
	pass := os.Getenv("PASS_GMAIL")
	if email == "" || pass == "" {
		log.Fatalln("configure sus variables de entorno")
		return
	}
	m := gomail.NewMessage()

	m.SetHeader("From", "luispfcanales@gmail.com")
	m.SetHeader("To", p.Email)

	m.SetHeader("Subject", "Gophers GO!")

	//t := template.Must(template.ParseFiles("./WellcomeTemplate.html"))
	//m.AddAlternativeWriter("text/html", func(w io.Writer) error {
	//	return t.Execute(w, "Registrate")
	//})

	m.SetBody("text/html", fmt.Sprintf("code verification: <b>%s</b>!", p.Code))

	d := gomail.NewDialer("smtp.gmail.com", 587, email, pass)

	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	w.Write([]byte(fmt.Sprintf("send email to: %s", p.Email)))
}

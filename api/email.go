package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	gomail "gopkg.in/mail.v2"
)

// Person is request model to send email
type Person struct {
	Code  string `json:"code,omitempty"`
	Email string `json:"email,omitempty"`
}

// Email send to email
func Email(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	t, err := template.ParseFiles("./WellcomeTemplate.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("not found template: %v", err)))
		return
	}
	t.Execute(&body, nil)
	//t := template.Must(template.ParseFiles("./WellcomeTemplate.html"))
	//m.AddAlternativeWriter("text/html", func(w io.Writer) error {
	//	return t.Execute(w, "Registrate")
	//})

	//m.SetBody("text/html", fmt.Sprintf("code verification: <b>%s</b>!", p.Code))
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, email, pass)

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	w.Write([]byte(fmt.Sprintf("send email to: %s", p.Email)))
}

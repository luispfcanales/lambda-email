package handler

import (
	"encoding/json"
	"fmt"
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

	m.SetBody("text/html", fmt.Sprintf("tu pedido esta en camino tu codigo es: <b>%s</b>!", p.Code))

	d := gomail.NewDialer("smtp.gmail.com", 587, email, pass)

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	w.Write([]byte(fmt.Sprintf("send email to: %s", p.Email)))
}

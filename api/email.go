package handler

import (
	"net/http"
	"os"
)

// Email send to email
func Email(w http.ResponseWriter, r *http.Request) {
	email := os.Getenv("email")
	pass := os.Getenv("email-pass")

	w.Write([]byte(email + pass + " => credentials"))
}

package handler

import (
	"net/http"
	"os"
)

// Message send to email
func Message(w http.ResponseWriter, r *http.Request) {
	email := os.Getenv("email")
	pass := os.Getenv("email-pass")

	w.Write([]byte(email + pass))
}

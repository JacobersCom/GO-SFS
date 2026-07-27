package main

import (
	"net/http"
)

type argon_params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	salt_length uint32
	key_length  uint32
}

func login(w http.ResponseWriter, r *http.Request) {

	params := &argon_params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		salt_length: 16,
		key_length:  32,
	}

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	check(err)

	username := r.FormValue("first")
	password := r.FormValue("password")

}

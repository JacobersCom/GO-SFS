package main

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"golang.org/x/crypto/argon2"
)

type argon_params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	salt_length uint32
	key_length  uint32
}

func GenerateRandomBytes(n uint32) ([]byte, error) {

	b := make([]byte, n)
	_, err := rand.Read(b)
	check(err)

	return b, nil
}

func GenerateFromPassword(password string, p *argon_params) (hash []byte, err error) {

	//generate a cryptographically secure random salt.
	salt, err := GenerateRandomBytes(p.salt_length)
	check(err)

	//generate the hash of the password using agron2id
	hash = argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.key_length)

	fmt.Println("Password hashed")

	return hash, nil
}

func Login(w http.ResponseWriter, r *http.Request) {

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

	password := r.FormValue("password")

	hash, err := GenerateFromPassword(password, params)
	check(err)

	fmt.Print(hash)

}

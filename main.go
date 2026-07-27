package main

import (
	"fmt"
	"net/http"
)

func server_start() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/login", Login)

	fmt.Println("Starting server at :4321")
	http.ListenAndServe(":4321", nil)
}

func main() {

	server_start()
}

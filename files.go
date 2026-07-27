package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Fprintf(w, "Error retrieving the file: %v", err)
		return
	}
	defer file.Close()

	f, err := os.Create("Files/" + handler.Filename)
	if err != nil {
		fmt.Fprintf(w, "Error opening the file: %v", err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
}

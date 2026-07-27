package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {

		panic(e)
	}
}

func make_dir(dir_name string) {

	err := os.Mkdir(dir_name, 0755)
	check(err)

}

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
	fmt.Fprintf(w, "File uploaded successfully: %v", handler.Filename)
}

func server_start() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/upload", uploadFile)
	fmt.Println("Starting server at :4321")
	http.ListenAndServe(":4321", nil)
}

func main() {

	server_start()
}

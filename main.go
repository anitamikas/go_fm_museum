package main

import (
	"fmt"
	"html/template"
	"net/http"

	"mikasanita.com/go/fm-museum/api"
	"mikasanita.com/go/fm-museum/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a Go program!"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.template")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}
	html.Execute(w, data.GetAll())

}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/hello", handleHello) // passing a function to the http handler, not invoking it
	server.HandleFunc("/templates", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)

	fs := http.FileServer(http.Dir("./public")) // create a file server, it automatically serves all files from public folder
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)
	if err == nil {
		fmt.Println("Error when opening the server")
	}
}

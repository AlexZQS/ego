package main

import (
	"html/template"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: ":80"}

	http.HandleFunc("/", welcome)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server.ListenAndServe()
}

func welcome(writer http.ResponseWriter, request *http.Request) {
	files, _ := template.ParseFiles("view/login.html")
	files.Execute(writer, nil)
}

package main

import (
	"ego/src/user"
	"html/template"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: ":8088"}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)

	//调用所有user模块的handler
	user.UserHandler()

	server.ListenAndServe()
}

func welcome(writer http.ResponseWriter, _ *http.Request) {
	files, _ := template.ParseFiles("view/login.html")
	_ = files.Execute(writer, nil)
}

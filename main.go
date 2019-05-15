package main

import (
	"ego/src/common"
	"ego/src/user"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func main() {

	//server := http.Server{
	//	Addr: ":8088"}
	//
	//mux.NewRouter()
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//http.HandleFunc("/", welcome)
	//
	////调用所有user模块的handler
	//user.UserHandler()
	//
	//server.ListenAndServe()

	common.Router.HandleFunc("/", welcome)
	common.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//满足/page/{page}格式的处理
	common.Router.HandleFunc("/page/{page}", showPage)
	user.UserHandler()
	_ = http.ListenAndServe(":80", common.Router)

}

func showPage(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	files, _ := template.ParseFiles("view/" + vars["page"] + ".html")
	_ = files.Execute(writer, nil)
}

func welcome(writer http.ResponseWriter, _ *http.Request) {
	files, _ := template.ParseFiles("view/login.html")
	_ = files.Execute(writer, nil)
}

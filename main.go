package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jessebarton/get/install"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	HandleError(w, err)

	v := req.FormValue("Images")
	install.Install(v)

}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

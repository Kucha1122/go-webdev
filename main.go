package main

import (

	"html/template"
	"log"
	"net/http"
)


var tpl *template.Template

type pageData struct {
	Title string
	FirstName string
}

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){
	log.Println("Server status [RUNNING]")
	http.HandleFunc("/", index)
	http.HandleFunc("/index", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, req *http.Request){

	pd := pageData{
		Title: "Index Page",
	}

	err := tpl.ExecuteTemplate(w,"index.gohtml",pd)

	if err != nil{
		log.Println(err)
		http.Error(w, "Internal server error!",http.StatusInternalServerError)
		return 
	}
}
func about(w http.ResponseWriter, req *http.Request){

	pd := pageData{
		Title: "About Page",
	}

	err := tpl.ExecuteTemplate(w,"about.gohtml",pd)

	if err != nil{
		log.Println(err)
		http.Error(w, "Internal server error!",http.StatusInternalServerError)
		return
	}
}
func contact(w http.ResponseWriter, req *http.Request){

	pd := pageData{
		Title: "Contact Page",
	}

	err := tpl.ExecuteTemplate(w,"contact.gohtml",pd)

	if err != nil{
		log.Println(err)
		http.Error(w, "Internal server error!",http.StatusInternalServerError)
		return
	}
}
func apply(w http.ResponseWriter, req *http.Request){

	pd := pageData{
		Title: "Index Page",
	}

	var firstName string
	if req.Method == http.MethodPost {
		firstName = req.FormValue("firstName")
		pd.FirstName = firstName

	}

	err := tpl.ExecuteTemplate(w,"apply.gohtml",pd)

	if err != nil{
		log.Println(err)
		http.Error(w, "Internal server error!",http.StatusInternalServerError)
		return
	}
}
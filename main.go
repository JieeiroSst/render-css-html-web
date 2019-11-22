package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var (
	linkDetails="./template/details.html"
	linkContact="./template/contact.html"
	linkIndex="./template/index.html"
	linkLogin="./template/login.html"
	linkRegister="./template/register.html"
	linkblog="./template/blog.html"
	linkProduct="./template/products.html"
)

func render(w http.ResponseWriter,data interface{},link string){
	tpl,err:=template.ParseFiles(link)
	if err!=nil{
		http.Error(w,"bug",http.StatusNotFound)
	}
	
	_ =tpl.Execute(w,data)
}

func renderAbout(w http.ResponseWriter, r* http.Request){
	render(w,nil,linkDetails)
}


func renderContact(w http.ResponseWriter,r * http.Request){
	render(w,nil,linkContact)
}


func renderIndex(w http.ResponseWriter,r *http.Request){
	render(w,nil,linkIndex)
}

func renderLogin(w http.ResponseWriter,r *http.Request){
	render(w,nil,linkLogin)
}

func renderRegister(w http.ResponseWriter,r *http.Request){
	render(w,nil,linkRegister)
}

func renderBlog(w http.ResponseWriter,r *http.Request){
	render(w,nil,linkblog)
}

func renderProduct(w http.ResponseWriter,r *http.Request){
	render(w,nil,linkProduct)
}

func main(){
	r:=mux.NewRouter()
	
	r.HandleFunc("/shop/details.html",renderAbout)
	r.HandleFunc("/shop/contact.html",renderContact)
	r.HandleFunc("/shop.com",renderIndex)
	r.HandleFunc("/shop/login.html",renderLogin)
	r.HandleFunc("/shop/register.html",renderRegister)
	r.HandleFunc("/shop/blog.html",renderBlog)
	r.HandleFunc("/shop/product.html",renderProduct)
	
	for i:=0;i < 8;i++{
		r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./template/css"))))
	}
	
	err:= http.ListenAndServe(":8080", r)
	if err!=nil{
		log.Fatal(err)
	}
}
package main
import (
	"net/http"
	"fmt"
	"html/template"
 )
 
 func main() {
 
	//We tell Go exactly where we can find our html file. We ask Go to parse the html file (Notice
	// the relative path). We wrap it in a call to template.Must() which handles any errors
	
	templates := template.Must(template.ParseGlob("template/*"))
 
	//Our HTML comes with CSS that Go needs to provide when we run the app. Here we tell go to create
	// a handle that looks in the static directory, go then uses the "/static/" as a url that our
	//html can refer to when looking for our css and other files. 
	
	http.Handle("/static/",
	   http.StripPrefix("/static/",
		  http.FileServer(http.Dir("static")))) //Go looks in the relative "static" directory first using http.FileServer(), then matches it to a
		  //url of our choice as shown in http.Handle("/static/"). This url is what we need when referencing our css files
		  //once the server begins. Our html code would therefore be <link rel="stylesheet"  href="/static/stylesheet/...">
		  //It is important to note the url in http.Handle can be whatever we like, so long as we are consistent.
 
	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
 
	   //If errors show an internal server error message
	   if err := templates.ExecuteTemplate(w, "homepage.html", nil); err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
	   }
	})
	fmt.Println("Listening");
	fmt.Println(http.ListenAndServe(":8080", nil));
 }
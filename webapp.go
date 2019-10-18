package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	

	fmt.Fprintf(w,`<h1>Go Lang</h1>
		<p1>Go is awesome </p1><br>
		<p1>.... and simple  language</p1>`)

	/*
	fmt.Fprintf(w, "<h1>Go Lang</h1>")
	fmt.Fprintf(w, "<p1>Go is awesome </p1><br>")
	fmt.Fprintf(w, "<p1>.... and simple  language</p1>")
	*/ 

}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About GoLang ")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil) 
}
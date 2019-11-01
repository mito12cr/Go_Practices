package main

import (
	"net/http"
)

func main() {

	//routes
	http.HandleFunc("/", homeHandler) //funcion que se encarga de manejar toda la funcionalidad
	http.HandleFunc("/contact", contactHandler)

	//start the server
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact Page"))
}

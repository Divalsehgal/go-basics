package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() error: %v", err), http.StatusInternalServerError)
		return
	} 
	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Printf("Name: %s, Address: %s\n", name, address)

	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello")
}

func main() {
	// Serve static files from the "static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

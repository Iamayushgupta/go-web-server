package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler handles the POST request to save form data.
func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() error: %v", err), http.StatusInternalServerError)
		log.Printf("Error parsing form: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %v\n", name)
	fmt.Fprintf(w, "Address: %v", address)
}

// helloHandler returns a simple "Hello" message.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		log.Printf("Method not supported: %v", r.Method)
		return
	}

	fmt.Fprintf(w, "Hello")
}

func main() {
	// Serve static files from the "./static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Register endpoint handlers
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	log.Println("Server started at port 8080")

	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(`
		<h1> Salut la classe! </h1>	
  `))
}

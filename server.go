package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "tipe.pdf")
	})

	log.Println("Server running on port 8888...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

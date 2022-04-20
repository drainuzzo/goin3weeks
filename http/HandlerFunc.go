package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greetings 2!")
	})
	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", h))
}

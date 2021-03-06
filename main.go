package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":9090", nil)
}

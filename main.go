package main

import (
	"PracticeProj/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "PracticeProject", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	http.Handle("/", hh)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Hi ")
	// })
	http.ListenAndServe(":8080", sm)
}

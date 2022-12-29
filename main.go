package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mhalsaiari/PracticeProject/handlers"
)

func main() {
	l := log.New(os.Stdout, "PracticeProj", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	http.Handle("/", hh)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Hi ")
	// })
	http.ListenAndServe(":8080", sm)
}

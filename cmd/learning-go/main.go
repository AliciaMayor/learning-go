package main

import (
	"log"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func main() {

	log.Println("Funciona")

	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request accepted!!")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("failed to start server: %v", err)
	}

}

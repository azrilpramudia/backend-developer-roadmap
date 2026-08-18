package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("Hello from Go"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("server running on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

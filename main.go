package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	select {
	case <-ctx.Done():
		log.Printf("context done: %v", ctx.Err())
	case <-time.After(5 * time.Second):
		log.Printf("Hello,  %s!", r.URL.Path[1:])
		_, err := fmt.Fprintf(w, "Hello, world!\n")
		if err != nil {
			log.Printf("err %v", err)
		}
		fmt.Fprintf(w, "Hostname: %s\n", host)
	}

}

func main() {
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	server := http.NewServeMux()
	server.HandleFunc("/", handler)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      server,
		Addr:         ":" + port,
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(srv.ListenAndServe())
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),

		ReadTimeout: 10 * time.Second,
		// WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second, // This specifies the keep-alive timeout
	}

	fmt.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

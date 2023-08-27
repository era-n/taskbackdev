package main

import (
	"log"
	"net/http"

	"github.com/era-n/taskbackdev/handlers"
)

func main() {
	http.HandleFunc("/auth", handlers.Authenticate)
	http.Handle("/refresh", handlers.RefreshMiddleware(http.HandlerFunc(handlers.Authenticate)))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err)
	}
}

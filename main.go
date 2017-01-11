package main

import (
	"fmt"
	"github.com/sailorgeoffrey/marathon-deployer/github"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT0")
	if port == "" {
		port = "8080"
	}
	http.Handle("/github/webhook", github.WebHookHandler{})
	http.Handle("/github/auth", github.AuthHandler{})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	http.HandleFunc("/__health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

package main

import (
	"fmt"
	"net/http"
)

func registerRoutes() {
	http.HandleFunc("/", hi)
	http.HandleFunc("/ping", ping)

	port := ":8080"
	fmt.Println("Server is listening at localhost" + port)
	http.ListenAndServe(port, nil)
}

func hi(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello, World!")
}

func ping(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "pong")
}

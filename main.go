package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	if name == "" {
		name = "world"
	}
	fmt.Printf("REQUEST: /%s\n", name)
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/", handler)

	port := 8080
	fmt.Printf("listening at %d ...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

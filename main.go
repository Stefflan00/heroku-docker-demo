package main

import (
	"fmt"
	"net/http"
	"os"
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
	port := os.Getenv("PORT")
	if port == "" {
		panic("need a $PORT env variable")
	}

	http.HandleFunc("/", handler)

	fmt.Printf("listening at :%s ...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

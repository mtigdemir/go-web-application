package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", handler)

	/* ListenAndServe always returns an error, since it only returns when an unexpected error occurs.
	In order to log that error we wrap the function call with log.Fatal. */
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Web Example Method: "+r.Method)
}

package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	body, _ := ioutil.ReadFile(title + ".txt")
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", title, string(body))
}

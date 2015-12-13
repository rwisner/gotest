package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fvbock/endless"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(endless.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Four")
}

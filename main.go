package main

import (
	"fmt"
	"log"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!!")
}

func main() {
	http.HandleFunc("/greet", greet)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

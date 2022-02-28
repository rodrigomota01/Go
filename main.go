package main

import (
	"fmt"
	"net/http"
)

func landing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Fala ai")
}

func main() {
	http.HandleFunc("/", landing)
	http.ListenAndServe(":2121", nil)
}

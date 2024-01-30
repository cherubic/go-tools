package main

import (
	"go-tools/function"
	"net/http"
)

func main() {
	http.HandleFunc("/", function.Home)
	http.HandleFunc("/url", function.GetUrlContent)

	http.ListenAndServe(":8080", nil)
}

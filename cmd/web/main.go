package main

import (
	"go-tools/function"
	"go-tools/page"
	"net/http"
)

func main() {
	router()

	http.ListenAndServe(":8080", nil)
}

func router() {
	http.HandleFunc("/", page.Home)

	http.HandleFunc("/url", function.GetURLContent)
	http.HandleFunc("/practice", function.Practice)

	http.HandleFunc("/fmtjson", page.FmtJSON)
	http.HandleFunc("/fmtjson/format", function.FmtJSON)
}

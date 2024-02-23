package page

import "net/http"

// Home is a function to handle home page
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("工具主页"))
}

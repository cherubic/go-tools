package page

import (
	"html/template"
	"net/http"
)

// FmtJSONPage is a struct to handle fmtjson.html
type FmtJSONPage struct {
	Title string
}

// FmtJSON is a function to format json
func FmtJSON(w http.ResponseWriter, r *http.Request) {
	p := &FmtJSONPage{}

	t, _ := template.ParseFiles("page/view/fmtjson.html")
	t.Execute(w, *p)
}

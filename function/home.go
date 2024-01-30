package function

import (
	"io"
	"net/http"
	"net/url"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("make some tool function with golang"))
}

func GetUrlContent(w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse(r.URL.Query().Get("url"))
	if err != nil {
		w.Write([]byte("url is empty"))
		return
	}

	res, err := http.Get(url.String())
	if err != nil {
		w.Write([]byte("get url content error"))
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		w.Write([]byte("get body error"))
	}

	w.Write(body)
}

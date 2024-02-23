package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetURLContent is a function to get url content
func GetURLContent(w http.ResponseWriter, r *http.Request) {
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

// FmtJSON is a function to format json
func FmtJSON(w http.ResponseWriter, r *http.Request) {
	type ReqBody struct {
		JSONStr string `json:"jsonStr"`
	}

	var reqBody ReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.Write([]byte("read body error"))
		return
	}

	jsonStr := reqBody.JSONStr

	var out bytes.Buffer
	b := []byte(jsonStr)
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		w.Write([]byte("format json error"))
	}

	w.Write(out.Bytes())
}

// Practice is a function to practice
func Practice(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	switch title {
	case "20240201":
		practice20240201()
	default:
		w.Write([]byte("title is empty"))
	}
}

func practice20240201() {
	i := 5
	defer practice20240201hello(i)
	i = i + 10
}

func practice20240201hello(i int) {
	fmt.Println(i)
}

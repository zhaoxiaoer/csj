package util

import (
	"bytes"
	"compress/gzip"
	"net/http"
	"strings"
)

func IsGzipSupported(r *http.Request) bool {
	encodings := r.Header.Get("Accept-Encoding")
	return strings.Contains(encodings, "gzip")
}

func IsGzipDisabled(r *http.Request) bool {
	r.ParseForm()
	flag, ok := r.Form["disableGzip"]
	return (ok && ((flag[0] == "") || (flag[0] == "true")))
}

func GetGzipData(s string) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte(s))
	w.Close()
	return b.Bytes()
}

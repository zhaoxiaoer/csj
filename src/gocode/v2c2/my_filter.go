package v2c2

import (
	"fmt"
	"mime"
	"net/http"
	"path"
	"time"
)

func init() {
	mime.AddExtensionType(".foo", "application/x-fubar")
}

func Reporter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 执行handler之前的逻辑
		// ...
		fmt.Printf("%s\n", r.RemoteAddr+" tried to access "+
			r.URL.RequestURI()+" on "+time.Now().Format(time.UnixDate+"."))

		// 前面的逻辑可能将w更改，此时就需要用新的w
		h.ServeHTTP(w, r)

		// 执行handler后的逻辑
		// ...
	})
}

func SetSpecialMIME(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置自定义MIME
		m := mime.TypeByExtension(path.Ext(r.URL.Path))
		if m != "" {
			w.Header().Set("Content-Type", m)
		}

		h.ServeHTTP(w, r)
	})
}

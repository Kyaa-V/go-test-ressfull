package middleware

import (
	"net/http"
	"time"
	"fmt"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start:= time.Now()
		next.ServeHTTP(w, r)

		duration:= time.Since(start)
		method:= r.Method
		path:= r.URL.Path
		fmt.Printf(
			"[LOGGER] %s %s %s %v\n",
			method,
			path,
			r.RemoteAddr,
			duration,
		)
	})
}
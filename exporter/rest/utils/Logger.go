package utils

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.HandlerFunc  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("Gend API: %s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
	})
}
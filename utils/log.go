package utils

import (
	"log"
	"net/http"
)

func LogRequest(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("%s request %s", req.RemoteAddr, req.URL)
		h.ServeHTTP(res, req)
	})
}

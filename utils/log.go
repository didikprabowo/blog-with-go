package utils

import (
	"log"
	"net/http"
)

func LogRequest(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf(" HOSTS : %s, REQUEST : %s, METHOD : %s\n", req.RemoteAddr, req.URL, req.Method)
		h.ServeHTTP(res, req)
	})
}

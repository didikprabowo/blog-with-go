package main

import (
	"flag"
	"github.com/didikprabowo/blog/cmd"
	"log"
	"net/http"
	"time"
)

func main() {
	r := cmd.AppRegister()

	var dir string

	flag.StringVar(&dir, "dir", "assets", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

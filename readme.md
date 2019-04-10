# CMS Blog With GO Language

## Requirement
1. Go Version : go1.11.4
2. MySQL Version : 5.7.25
## Package
1. MySQL - github.com/go-sql-driver/mysql
2. MUX - github.com/gorilla/mux
3. Session Gorilla - github.com/gorilla/sessions
4. Bcrypt - golang.org/x/crypto

## How to use
1. go get github.com/didikprabowo/blog
3. Create main.go
```
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
		Addr:         "127.0.0.1:8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

```
3. Run 
```
go run main.go
```
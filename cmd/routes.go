package cmd

import (
	"github.com/didikprabowo/blog/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Method  string
	Handler http.HandlerFunc
	Path    string
}

func AppRegister() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	routes := DataRoutes()
	for _, v := range routes {
		r.Path(v.Path).HandlerFunc(v.Handler).Methods(v.Method)
	}
	return r
}

func DataRoutes() []Route {
	routes := []Route{
		Route{
			Method:  "GET",
			Handler: Welcome,
			Path:    "/welcome",
		},
		Route{
			Method:  "GET",
			Handler: handlers.Auth,
			Path:    "/auth",
		},
		Route{
			Method:  "POST",
			Handler: handlers.Login,
			Path:    "/login",
		},
		Route{
			Method:  "GET",
			Handler: handlers.Register,
			Path:    "/register",
		},
		Route{
			Method:  "GET",
			Handler: handlers.Logout,
			Path:    "/logout",
		},
	}
	return routes
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome My Blog .."))
}
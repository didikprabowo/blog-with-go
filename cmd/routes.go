package cmd

import (
	"fmt"
	"github.com/didikprabowo/blog/handlers"
	"github.com/didikprabowo/blog/handlers/admin"
	"github.com/didikprabowo/blog/handlers/web"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
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

var store = sessions.NewCookieStore([]byte("didikprabowo"))

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "login")
		if err != nil {
			panic(err.Error)
		}
		if len(session.Values) == 0 {
			http.Redirect(w, r, "/auth", 301)
		} else {
			next.ServeHTTP(w, r)
		}
		fmt.Println(len(session.Values))
	})
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
		Route{
			Method:  "GET",
			Handler: loggingMiddleware(admin.GetCategory),
			Path:    "/admin/category",
		},
		Route{
			Method:  "GET",
			Handler: loggingMiddleware(admin.CreateCategory),
			Path:    "/admin/category/create",
		},
		Route{
			Method:  "POST",
			Handler: loggingMiddleware(admin.StoreCategory),
			Path:    "/admin/category/store",
		},
		Route{
			Method:  "GET",
			Handler: loggingMiddleware(admin.EditCategory),
			Path:    "/admin/category/edit/{id}",
		},
		Route{
			Method:  "POST",
			Handler: loggingMiddleware(admin.UpdateCategory),
			Path:    "/admin/category/update",
		},
		Route{
			Method:  "GET",
			Handler: loggingMiddleware(admin.DeleteCategory),
			Path:    "/admin/category/delete/{id}",
		},
		Route{
			Method:  "GET",
			Handler: loggingMiddleware(admin.GetPosts),
			Path:    "/admin/post",
		},
		Route{
			Method:  "GET",
			Handler: loggingMiddleware(admin.CreatePost),
			Path:    "/admin/post/create",
		},
		Route{
			Method:  "POST",
			Handler: loggingMiddleware(admin.StorePost),
			Path:    "/admin/post/store",
		},
		Route{
			Method:  "GET",
			Handler: loggingMiddleware(admin.EditPost),
			Path:    "/admin/post/edit/{id}",
		},
		Route{
			Method:  "POST",
			Handler: loggingMiddleware(admin.UpdatePost),
			Path:    "/admin/post/update",
		},
		Route{
			Method:  "GET",
			Handler: loggingMiddleware(admin.DeletePost),
			Path:    "/admin/post/delete/{id}",
		},
		Route{
			Method:  "GET",
			Handler: web.Beranda,
			Path:    "/",
		},
		Route{
			Method:  "GET",
			Handler: web.DetailPosts,
			Path:    "/{slug}",
		},
	}
	return routes
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome My Blog .."))
}

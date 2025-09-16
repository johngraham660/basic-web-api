package router

import (
	"net/http"

	"basic_web_api/auth"
	"basic_web_api/handlers"
)

func NewRouter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/register":
			handlers.RegisterHandler(w, r)
		case "/login":
			handlers.LoginHandler(w, r)
		case "/users":
			auth.AuthMiddleware(handlers.UsersHandler).ServeHTTP(w, r)
		case "/posts":
			auth.AuthMiddleware(handlers.PostsHandler).ServeHTTP(w, r)
		default:
			handlers.NotFoundHandler(w, r)
		}
	}
}

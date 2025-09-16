// Package main Basic Web API
//
// @title Basic Web API
// @version 1.0
// @description A simple RESTful API built with Go and PostgreSQL.
// @host localhost:8080
// @BasePath /
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"basic_web_api/db"
	_ "basic_web_api/docs"
	"basic_web_api/router"

	httpSwagger "github.com/swaggo/http-swagger"
)

func serveTemplate(tmpl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/" + tmpl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	}
}

func main() {
	// Initialize database
	if err := db.InitDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Web routes
	http.HandleFunc("/login", serveTemplate("login.html"))
	http.HandleFunc("/register", serveTemplate("register.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		http.NotFound(w, r)
	})

	// API routes
	apiHandler := router.NewRouter()
	http.Handle("/api/", http.StripPrefix("/api", apiHandler))

	// Swagger UI route
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Server is running on port 8080")
	fmt.Println("View API documentation at http://localhost:8080/swagger/index.html")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

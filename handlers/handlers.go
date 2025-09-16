// Package handlers provides HTTP request handlers for the API
package handlers

import (
	"encoding/json"
	"net/http"

	"basic_web_api/models"
	"basic_web_api/repository"
)

var (
	userRepo = repository.NewUserRepository()
	postRepo = repository.NewPostRepository()
)

// UsersHandler godoc
// @Summary Handle users
// @Description Get all users or create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User false "User object for creation"
// @Success 200 {array} models.User
// @Success 201 {object} models.User
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /users [get]
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := userRepo.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := userRepo.Create(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// PostsHandler godoc
// @Summary Handle posts
// @Description Get all posts or create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body models.Post false "Post object for creation"
// @Success 200 {array} models.Post
// @Success 201 {object} models.Post
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /posts [get]
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		posts, err := postRepo.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)

	case http.MethodPost:
		var post models.Post
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := postRepo.Create(&post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// NotFoundHandler godoc
// @Summary Handle 404 errors
// @Description Returns a 404 Not Found error
// @Tags errors
// @Produce json
// @Success 404 {object} string
// @Router /404 [get]
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "404 Not Found", http.StatusNotFound)
}

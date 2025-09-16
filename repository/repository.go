package repository

import (
	"basic_web_api/db"
	"basic_web_api/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := db.DB.Select(&users, "SELECT id, name, dob FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (name, email, password, dob) VALUES ($1, $2, $3, $4) RETURNING id`
	return db.DB.QueryRow(query, user.Name, user.Email, user.Password, user.DoB).Scan(&user.ID)
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password, dob FROM users WHERE email = $1`
	err := db.DB.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type PostRepository struct{}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	var posts []models.Post
	err := db.DB.Select(&posts, "SELECT id, content, created_at FROM posts")
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) Create(post *models.Post) error {
	query := `INSERT INTO posts (content) VALUES ($1) RETURNING id, created_at`
	return db.DB.QueryRow(query, post.Content).Scan(&post.ID, &post.CreatedAt)
}

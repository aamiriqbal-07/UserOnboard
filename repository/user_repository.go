package repository

import (
	"database/sql"
	"log"
	"userOnboard/config"
	"userOnboard/models"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByID(id string) (*models.User, error)
	ListUsers() ([]models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{db: config.DB}
}

func (r *userRepository) CreateUser(user models.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name, signup_time) VALUES (?, ?, ?)", user.ID, user.Name, user.SignupTime)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
	}
	return err
}

func (r *userRepository) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, name, signup_time FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.SignupTime)
	if err != nil {
		log.Printf("Failed to get user by ID: %v", err)
		return nil, err
	}
	return user, nil
}

func (r *userRepository) ListUsers() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name, signup_time FROM users")
	if err != nil {
		log.Printf("Failed to list users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.SignupTime); err != nil {
			log.Printf("Failed to scan user: %v", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

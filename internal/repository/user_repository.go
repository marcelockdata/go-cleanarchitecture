package repository

import (
	"cleanarchitecture/internal/domain"
	"database/sql"
	"log"
)

type UserRepository interface {
	FindByUsername(username string) (*domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, username, password FROM users WHERE username = ?"
	row := r.db.QueryRow(query, username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User not found: %s", username) // Log de depuração
			return nil, domain.ErrInvalidCredentials
		}
		log.Printf("Database error: %v", err) // Log de erro
		return nil, err
	}
	return &user, nil
}

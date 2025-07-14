package handler

import (
	"cleanarchitecture/internal/usecase"
	"encoding/json"
	"log"
	"net/http"
)

// Credentials represents the user credentials for authentication.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

// Authenticate godoc
// @Summary Authenticate a user
// @Description Authenticate a user and return a token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param credentials body Credentials true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth [post]
func (h *UserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Printf("Error decoding request body: %v", err) // Log de erro
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Authenticating user: %s", credentials.Username) // Log de depuração

	authenticated, err := h.userUsecase.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		log.Printf("Authentication error: %v", err) // Log de erro
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if !authenticated {
		log.Printf("Invalid credentials for user: %s", credentials.Username) // Log de depuração
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := h.userUsecase.GenerateToken(credentials.Username)
	if err != nil {
		log.Printf("Token generation error: %v", err) // Log de erro
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

package routes

import (
	"encoding/json"
	"log/slog"
	"net/http"
	sqlite "rolerocket/internal/db"
	"rolerocket/internal/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSearch struct {
	Username string
}

func GetUsers(w http.ResponseWriter, r *http.Request) ([]string, error) {
	userSearch := r.URL.Query().Get("username")
	return sqlite.DBInstance.GetUsers(r.Context(), &userSearch)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var creds Credentials

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid or unexpected JSON fields", http.StatusBadRequest)
		return
	}

	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Missing username or password in JSON body", http.StatusBadRequest)
		return
	}

	if len(creds.Username) <= 5 || len(creds.Password) <= 5 {
		http.Error(w, "Username and password has to be longer than 5 characters", http.StatusBadRequest)
		return
	}

	existingUsers, err := sqlite.DBInstance.GetUsers(r.Context(), &creds.Username)
	if err != nil {
		logger.Error(r.Context(), "Error thrown whilst getting user", slog.Any("error", err))
		http.Error(w, "Missing username or password in JSON body", http.StatusInternalServerError)
		return
	}
	if len(existingUsers) != 0 {
		http.Error(w, "Username already exists!", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(r.Context(), "Error thrown whilst salting and hashing password", slog.Any("error", err))
		http.Error(w, "Internal error while securing password", http.StatusInternalServerError)
		return
	}
	sqlite.DBInstance.InsertUser(r.Context(), creds.Username, string(hashedPassword))
}

func UpdateUser() {

}

func DeleteUser() {

}

func GetToken(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid or unexpected JSON fields", http.StatusBadRequest)
		return
	}

	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Missing username or password in JSON body", http.StatusBadRequest)
		return
	}

	err = sqlite.DBInstance.VerifyLogin(r.Context(), &creds.Username, &creds.Password)
	if err != nil {
		http.Error(w, "Username or password is wrong", http.StatusUnauthorized)
		return
	}

	secretKey := []byte("your-secret-key")
	claims := jwt.MapClaims{
		"name": creds.Username,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		logger.Error(r.Context(), "Error signing token:", slog.Any("error", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": signedToken,
	})
}

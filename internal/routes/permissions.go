package routes

import (
	"encoding/json"
	"log/slog"
	"net/http"
	sqlite "rolerocket/internal/db"
	"rolerocket/internal/logger"
)

type Permission struct {
	Name string `json:"name"`
}

type Role struct {
	Name string `json:"name"`
}

func getPermissions(w http.ResponseWriter, r *http.Request) {
	permissionSearch := r.URL.Query().Get("name")
	permissions, err := sqlite.DBInstance.GetPermissions(r.Context(), permissionSearch)
	if err != nil {
		respondWithError(w, "Failed to fetch permissions", http.StatusInternalServerError)
		logger.Error(r.Context(), "Failed to fetch permissions:", slog.Any("error", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(permissions); err != nil {
		respondWithError(w, "Failed to encode permissions", http.StatusInternalServerError)
		logger.Error(r.Context(), "Failed to encode permissions", slog.Any("error", err))
		return
	}
}

func insertPermission(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var permission Permission

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&permission)
	if err != nil {
		respondWithError(w, "Invalid or unexpected JSON fields", http.StatusBadRequest)
		return
	}

	if permission.Name == "" {
		respondWithError(w, "Missing name in JSON body", http.StatusBadRequest)
		return
	}

	existingPermissions, err := sqlite.DBInstance.GetPermissions(r.Context(), permission.Name)
	if err != nil {
		logger.Error(r.Context(), "Error thrown whilst getting existing permissions", slog.Any("error", err))
		respondWithError(w, "Error thrown whilst getting existing permissions", http.StatusInternalServerError)
		return
	}
	if len(existingPermissions) != 0 {
		respondWithError(w, "Permission name already exists!", http.StatusBadRequest)
		return
	}

	err = sqlite.DBInstance.InsertPermission(r.Context(), permission.Name)
	if err != nil {
		respondWithError(w, "Error whilst inserting permission", http.StatusBadRequest)
		logger.Error(r.Context(), "Error thrown whilst inserting permission", slog.Any("error", err))
		return
	}
}

func getRoles(w http.ResponseWriter, r *http.Request) {
	roleSearch := r.URL.Query().Get("name")
	roles, err := sqlite.DBInstance.GetRoles(r.Context(), roleSearch)
	if err != nil {
		respondWithError(w, "Failed to fetch roles", http.StatusUnauthorized)
		logger.Error(r.Context(), "Failed to fetch roles:", slog.Any("error", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(roles); err != nil {
		respondWithError(w, "Failed to encode roles", http.StatusInternalServerError)
		logger.Error(r.Context(), "Failed to encode roles", slog.Any("error", err))
		return
	}
}

func insertRole(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var role Role

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&role)
	if err != nil {
		respondWithError(w, "Invalid or unexpected JSON fields", http.StatusBadRequest)
		return
	}

	if role.Name == "" {
		respondWithError(w, "Missing name in JSON body", http.StatusBadRequest)
		return
	}

	existingRole, err := sqlite.DBInstance.GetRoles(r.Context(), role.Name)
	if err != nil {
		logger.Error(r.Context(), "Error thrown whilst getting existing roles", slog.Any("error", err))
		respondWithError(w, "Error thrown whilst getting existing roles", http.StatusInternalServerError)
		return
	}
	if len(existingRole) != 0 {
		respondWithError(w, "Role name already exists!", http.StatusBadRequest)
		return
	}

	err = sqlite.DBInstance.InsertRole(r.Context(), role.Name)
	if err != nil {
		respondWithError(w, "Error whilst inserting role", http.StatusBadRequest)
		logger.Error(r.Context(), "Error thrown whilst inserting role", slog.Any("error", err))
		return
	}
}

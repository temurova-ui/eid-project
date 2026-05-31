package handlers

import (
	"encoding/json"
	"net/http"
	"project/logger"
	"project/models"
	"project/storage"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type UserHandler struct {
	Storage *storage.UserStorage
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Storage.GetAll()
	if err != nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return 
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.Storage.GetByID(id)
	if err != nil{
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		logger.L.Error(
			"failed to decode user",
			zap.Error(err),
		)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Storage.Create(user)
	if err != nil{
		logger.L.Error(
			"failed to create user",
			zap.Error(err),
		)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Storage.Update(id, user)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sayed-imran/go-design-pattern/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (handlerRepo *Repository) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.APIUser
	_ = json.NewDecoder(r.Body).Decode(&user)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	id := primitive.NewObjectID()
	handlerRepo.DB.AddUser(ctx, models.User{
		ID:        id,
		UserName:  user.UserName,
		FisrtName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	json.NewEncoder(w).Encode(user)
}

func (handlerRepo *Repository) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	id := r.URL.Query().Get("id")
	user, err := handlerRepo.DB.FindSingleUser(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (handlerRepo *Repository) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.APIUser
	_ = json.NewDecoder(r.Body).Decode(&user)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	id := r.URL.Query().Get("id")
	handlerRepo.DB.UpdateUser(ctx, id, models.User{
		UserName:  user.UserName,
		FisrtName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	json.NewEncoder(w).Encode(user)
}

func (handlerRepo *Repository) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	id := r.URL.Query().Get("id")
	handlerRepo.DB.DeleteUser(ctx, id)
	w.WriteHeader(http.StatusOK)
}

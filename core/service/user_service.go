package service

import (
	"context"
	"net/http"
	"time"

	"github.com/ertugrul-k/goap/core/repo"
	. "github.com/ertugrul-k/goap/db"
	"github.com/gorilla/mux"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("content-type", "application/json")
	coll := DB.Database.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	users := repo.FindAll(ctx, coll)
	RespondWithJson(w, http.StatusCreated, "", users)
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("content-type", "application/json")
	coll := DB.Database.Collection("users")
	user_id := mux.Vars(r)["_id"]
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	user := repo.FindOne(ctx, coll, user_id)
	RespondWithJson(w, http.StatusCreated, "", user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	coll := DB.Database.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	user := repo.CreateUser(ctx, coll, r.Body)

	RespondWithJson(w, http.StatusCreated, "", user)
}

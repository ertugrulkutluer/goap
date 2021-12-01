package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ertugrul-k/goap/core/repo"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAll(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("content-type", "application/json")
	coll := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	users := repo.FindAll(ctx, coll)
	json.NewEncoder(w).Encode(users)
}

func FindOne(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("content-type", "application/json")
	coll := db.Collection("users")
	user_id := mux.Vars(r)["_id"]
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	user := repo.FindOne(ctx, coll, user_id)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	w.Header().Set("Content-Type", "application/json")
	coll := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	user := repo.CreateUser(ctx, coll, r.Body)

	respondWithJson(w, http.StatusCreated, "", user)
}

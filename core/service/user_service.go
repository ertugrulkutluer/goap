package service

import "github.com/ertugrul-k/goap/models"

type UserService struct {
}

func CreateUser(ucr request.UserCreateRequest) *models.User

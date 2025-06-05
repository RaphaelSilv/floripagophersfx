package repository

import (
	"fmt"
	"github.com/raphaelsilv/fgophers/example2/models"
)

type UserRetriever interface {
	GetUser(string) models.User
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) GetUser(userID string) models.User {
	fmt.Printf("repository: getting user with ID: %s\n", userID)
	return models.User{Name: "Gopher"}
}

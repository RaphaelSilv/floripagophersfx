package repository

import (
	"fmt"
	"github.com/raphaelsilv/fgophers/example5/compliancefx/models"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewUserRepository)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) GetUser(userID string) models.User {
	fmt.Printf("repository: getting user with ID: %s\n", userID)
	return models.User{Name: "Gopher"}
}

package client

import (
	"fmt"
	"github.com/raphaelsilv/fgophers/example4/models"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewUserClient)

type UserClient struct{}

func NewUserClient() *UserClient {
	return &UserClient{}
}

func (u UserClient) GetUser(userID string) models.User {
	fmt.Printf("client: getting user with ID: %s\n", userID)
	return models.User{Name: "Gopher"}
}

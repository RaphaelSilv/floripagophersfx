package client

import (
	"fmt"
	"github.com/raphaelsilv/fgophers/example2/models"
)

type UserClient struct{}

func NewUserClient() *UserClient {
	return &UserClient{}
}

func (u UserClient) GetUser(userID string) models.User {
	fmt.Printf("client: getting user with ID: %s\n", userID)
	return models.User{Name: "Gopher"}
}

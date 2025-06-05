package service

import (
	"fmt"
	"github.com/raphaelsilv/fgophers/example2/models"
	"github.com/raphaelsilv/fgophers/example3/handler"
)

type UserRetriever interface {
	GetUser(string) models.User
}

type Service struct {
	client UserRetriever
}

func NewService(client UserRetriever) *Service {
	return &Service{client: client}
}

func New(client UserRetriever) handler.Service {
	return &Service{client: client}
}

func (s *Service) Checkout(userID string) {
	user := s.client.GetUser(userID)
	fmt.Printf("service: checked out user: %s\n", user.Name)
}

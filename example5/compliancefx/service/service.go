package service

import (
	"fmt"
	"github.com/raphaelsilv/fgophers/example4/client"
	"github.com/raphaelsilv/fgophers/example4/handler"
	"github.com/raphaelsilv/fgophers/example4/models"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewService,
		fx.As(new(handler.Service)),
		fx.From(new(*client.UserClient)),
	),
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

func (s *Service) Checkout(userID string) {
	user := s.client.GetUser(userID)
	fmt.Printf("service: checked out user: %s\n", user.Name)
}

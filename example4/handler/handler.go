package handler

import "go.uber.org/fx"

var module = fx.Provide(
	NewHandler,
)

type Service interface {
	Checkout(userID string)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) Process(userID string) {
	h.service.Checkout(userID)
}

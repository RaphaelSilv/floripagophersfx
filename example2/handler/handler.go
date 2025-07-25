package handler

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

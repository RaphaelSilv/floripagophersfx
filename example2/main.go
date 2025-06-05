package main

import (
	"github.com/raphaelsilv/fgophers/example2/handler"
	"github.com/raphaelsilv/fgophers/example2/repository"
	"github.com/raphaelsilv/fgophers/example2/service"
)

func main() {
	userRetriever := repository.NewUserRepository()
	//userRetriever := client.NewUserClient()
	service := service.NewService(userRetriever)
	handler := handler.NewHandler(service)

	handler.Process("1234")
}

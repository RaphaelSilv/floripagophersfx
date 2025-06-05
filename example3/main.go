package main

import (
	"github.com/raphaelsilv/fgophers/example3/handler"
	"github.com/raphaelsilv/fgophers/example3/repository"
	"github.com/raphaelsilv/fgophers/example3/service"
	"go.uber.org/fx"
)

// 	"go.uber.org/fx"

func main() {
	app := fx.New(
		fx.Provide(
			//service.NewService,
			//	repository.NewUserRepository,
			//	client.NewUserClient,
			handler.NewHandler,
		),
		// Start HTTP Server
		fx.Invoke(func(handler *handler.Handler) {
			handler.Process("123")
		}),
	)
	app.Run()
}

// LOREM IPSUM
// LOREM IPSUM
// LOREM IPSUM
// LOREM IPSUM
// LOREM IPSUM

// Here we solve two problems. The first with Fx.As
// since it binds the NewService construct with the actual interface in need i.e. handler.Service.
// The second problem lies in the constructor of the service as it depends on an ambiguous interface;
func annotate() interface{} {
	return fx.Annotate(
		service.NewService,
		fx.As(new(handler.Service)),
		fx.From(new(*repository.UserRepository)),
	)
}

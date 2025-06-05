package main

import (
	"github.com/raphaelsilv/fgophers/example4/handler"
	"github.com/raphaelsilv/fgophers/example4/internal"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		internal.Module,
		fx.Invoke(func(handler *handler.Handler) {
			handler.Process("123")
		}),
	)
	app.Run()
}

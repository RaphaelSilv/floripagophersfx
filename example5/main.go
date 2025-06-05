package main

import (
	"github.com/raphaelsilv/fgophers/example5/compliancefx"
	"github.com/raphaelsilv/fgophers/example5/compliancefx/handler"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		compliancefx.Module,
		compliancefx.Module,
		compliancefx.Module,
		fx.Invoke(func(handler *handler.Handler) {
			handler.Process("123")
		}),
	)
	app.Run()
}

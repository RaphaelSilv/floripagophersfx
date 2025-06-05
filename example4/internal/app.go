package internal

import (
	"github.com/raphaelsilv/fgophers/example4/client"
	"github.com/raphaelsilv/fgophers/example4/handler"
	"github.com/raphaelsilv/fgophers/example4/repository"
	"github.com/raphaelsilv/fgophers/example4/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handler.Module,
	service.Module,
	repository.Module,
	client.Module,
)

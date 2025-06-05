package compliancefx

import (
	"github.com/raphaelsilv/fgophers/example5/compliancefx/client"
	"github.com/raphaelsilv/fgophers/example5/compliancefx/handler"
	"github.com/raphaelsilv/fgophers/example5/compliancefx/repository"
	"github.com/raphaelsilv/fgophers/example5/compliancefx/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handler.Module,
	service.Module,
	repository.Module,
	client.Module,
)

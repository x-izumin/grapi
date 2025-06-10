package svcgen

import (
	"github.com/x-izumin/grapi/pkg/gencmd"
	"github.com/x-izumin/grapi/pkg/protoc"
	"github.com/x-izumin/grapi/pkg/svcgen/params"
)

type CreateAppFunc func(*gencmd.Command) (*App, error)

type App struct {
	ProtocWrapper protoc.Wrapper
	ParamsBuilder params.Builder
}

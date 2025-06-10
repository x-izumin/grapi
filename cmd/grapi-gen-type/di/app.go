package di

import (
	"github.com/x-izumin/grapi/pkg/gencmd"
	"github.com/x-izumin/grapi/pkg/protoc"
)

type CreateAppFunc func(*gencmd.Command) (*App, error)

type App struct {
	Protoc protoc.Wrapper
}

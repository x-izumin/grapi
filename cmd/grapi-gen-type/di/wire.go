//+build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/x-izumin/grapi/pkg/cli"
	"github.com/x-izumin/grapi/pkg/gencmd"
	"github.com/x-izumin/grapi/pkg/protoc"
)

func NewApp(*gencmd.Command) (*App, error) {
	wire.Build(
		App{},
		gencmd.Set,
		cli.UIInstance,
		protoc.WrapperSet,
	)
	return nil, nil
}

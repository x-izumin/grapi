//+build wireinject

package testing

import (
	"github.com/google/wire"

	"github.com/x-izumin/grapi/pkg/cli"
	"github.com/x-izumin/grapi/pkg/gencmd"
	"github.com/x-izumin/grapi/pkg/protoc"
	"github.com/x-izumin/grapi/pkg/svcgen"
)

func NewTestApp(*gencmd.Command, protoc.Wrapper, cli.UI) (*svcgen.App, error) {
	wire.Build(
		gencmd.Set,
		svcgen.ProvideParamsBuilder,
		svcgen.App{},
	)
	return nil, nil
}

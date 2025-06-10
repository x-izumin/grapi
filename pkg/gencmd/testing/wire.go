//go:build wireinject
// +build wireinject

package testing

import (
	"github.com/google/wire"

	"github.com/x-izumin/grapi/pkg/cli"
	"github.com/x-izumin/grapi/pkg/gencmd"
)

func NewTestApp(*gencmd.Command, cli.UI) (*gencmd.App, error) {
	wire.Build(
		gencmd.Set,
	)
	return nil, nil
}

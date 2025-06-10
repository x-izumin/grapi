package svcgen

import (
	"github.com/google/wire"

	"github.com/x-izumin/grapi/pkg/cli"
	"github.com/x-izumin/grapi/pkg/grapicmd"
	"github.com/x-izumin/grapi/pkg/protoc"
	"github.com/x-izumin/grapi/pkg/svcgen/params"
	_ "github.com/x-izumin/grapi/pkg/svcgen/template"
)

func ProvideParamsBuilder(rootDir cli.RootDir, protocCfg *protoc.Config, grapiCfg *grapicmd.Config) params.Builder {
	return params.NewBuilder(
		rootDir,
		protocCfg.ProtosDir,
		protocCfg.OutDir,
		grapiCfg.Grapi.ServerDir,
		grapiCfg.Package,
	)
}

var Set = wire.NewSet(
	ProvideParamsBuilder,
	App{},
)

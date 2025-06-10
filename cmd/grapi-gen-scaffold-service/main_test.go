package main

import (
	"context"
	"testing"

	"github.com/spf13/afero"

	"github.com/izumin5210/clig/pkg/clib"
	"github.com/x-izumin/grapi/pkg/cli"
	"github.com/x-izumin/grapi/pkg/gencmd"
	gencmdtesting "github.com/x-izumin/grapi/pkg/gencmd/testing"
	"github.com/x-izumin/grapi/pkg/grapicmd"
	"github.com/x-izumin/grapi/pkg/protoc"
	"github.com/x-izumin/grapi/pkg/svcgen"
	svcgentesting "github.com/x-izumin/grapi/pkg/svcgen/testing"
)

func TestRun(t *testing.T) {
	cases := []svcgentesting.Case{
		{
			Test:  "simple",
			GArgs: []string{"book"},
			DArgs: []string{"book"},
			Files: []string{
				"api/protos/book.proto",
				"app/server/book_server.go",
				"app/server/book_server_register_funcs.go",
				"app/server/book_server_test.go",
			},
		},
		{
			Test:  "specify resource name",
			GArgs: []string{"library", "--resource-name=book"},
			DArgs: []string{"library"},
			Files: []string{
				"api/protos/library.proto",
				"app/server/library_server.go",
				"app/server/library_server_register_funcs.go",
				"app/server/library_server_test.go",
			},
		},
	}

	rootDir := cli.RootDir{clib.Path("/home/src/testapp")}

	createSvcApp := func(cmd *gencmd.Command) (*svcgen.App, error) {
		return svcgentesting.NewTestApp(cmd, &fakeProtocWrapper{}, cli.NopUI)
	}
	createGenApp := func(cmd *gencmd.Command) (*gencmd.App, error) {
		return gencmdtesting.NewTestApp(cmd, cli.NopUI)
	}
	createCmd := func(t *testing.T, fs afero.Fs, tc svcgentesting.Case) gencmd.Executor {
		ctx := &grapicmd.Ctx{
			FS:      fs,
			RootDir: rootDir,
			Config: grapicmd.Config{
				Package: tc.PkgName,
			},
			ProtocConfig: protoc.Config{
				ProtosDir: tc.ProtoDir,
				OutDir:    tc.ProtoOutDir,
			},
		}
		ctx.Config.Grapi.ServerDir = tc.ServerDir
		return buildCommand(createSvcApp, gencmd.WithGrapiCtx(ctx), gencmd.WithCreateAppFunc(createGenApp))
	}

	ctx := &svcgentesting.Ctx{
		GOPATH:    "/home",
		RootDir:   rootDir,
		CreateCmd: createCmd,
		Cases:     cases,
	}

	svcgentesting.Run(t, ctx)
}

type fakeProtocWrapper struct{}

func (*fakeProtocWrapper) Exec(context.Context) error { return nil }

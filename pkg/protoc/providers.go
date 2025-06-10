package protoc

import (
	"sync"

	"github.com/google/wire"
	"github.com/izumin5210/clig/pkg/clib"
	"github.com/izumin5210/execx"
	"github.com/izumin5210/gex"
	"github.com/izumin5210/gex/pkg/tool"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"go.uber.org/zap"

	"github.com/x-izumin/grapi/pkg/cli"
)

var (
	gexCfg   *gex.Config
	gexCfgMu sync.Mutex

	toolRepo   tool.Repository
	toolRepoMu sync.Mutex
)

func ProvideGexConfig(
	fs afero.Fs,
	exec *execx.Executor,
	io *clib.IO,
	rootDir cli.RootDir,
) *gex.Config {
	gexCfgMu.Lock()
	defer gexCfgMu.Unlock()
	if gexCfg == nil {
		gexCfg = &gex.Config{
			OutWriter:  io.Out,
			ErrWriter:  io.Err,
			InReader:   io.In,
			FS:         fs,
			Exec:       exec,
			WorkingDir: rootDir.String(),
			Verbose:    clib.IsVerbose() || clib.IsDebug(),
			Logger:     zap.NewStdLog(zap.L()),
		}
	}
	return gexCfg
}

func ProvideToolRepository(gexCfg *gex.Config) (tool.Repository, error) {
	toolRepoMu.Lock()
	defer toolRepoMu.Unlock()
	if toolRepo == nil {
		var err error
		toolRepo, err = gexCfg.Create()
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return toolRepo, nil
}

// WrapperSet is a provider set that includes gex things and Wrapper instance.
var WrapperSet = wire.NewSet(
	ProvideGexConfig,
	ProvideToolRepository,
	NewWrapper,
)

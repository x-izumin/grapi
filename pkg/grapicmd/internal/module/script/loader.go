package script

import (
	"path/filepath"
	"runtime"
	"sort"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"go.uber.org/zap"

	"github.com/izumin5210/clig/pkg/clib"
	"github.com/izumin5210/execx"
	"github.com/x-izumin/grapi/pkg/grapicmd/internal/module"
	"github.com/x-izumin/grapi/pkg/grapicmd/util/fs"
)

// NewLoader creates a new ScriptLoader instance.
func NewLoader(fs afero.Fs, io *clib.IO, exec *execx.Executor, rootDir string) module.ScriptLoader {
	return &scriptLoader{
		fs:      fs,
		io:      io,
		exec:    exec,
		rootDir: rootDir,
		binDir:  filepath.Join(rootDir, "bin"),
		scripts: make(map[string]module.Script),
	}
}

type scriptLoader struct {
	fs      afero.Fs
	io      *clib.IO
	exec    *execx.Executor
	rootDir string
	binDir  string
	scripts map[string]module.Script
	names   []string
}

func (f *scriptLoader) Load(dir string) error {
	srcsByDir, err := fs.FindMainPackagesAndSources(f.fs, dir)
	zap.L().Debug("found main packages", zap.Any("srcs_by_dir", srcsByDir))
	if err != nil {
		return errors.Wrap(err, "failed to find commands")
	}
	for dir, srcs := range srcsByDir {
		srcPaths := make([]string, 0, len(srcs))
		for _, name := range srcs {
			srcPaths = append(srcPaths, filepath.Join(dir, name))
		}
		name := filepath.Base(dir)
		ext := ""
		if runtime.GOOS == "windows" {
			ext = ".exe"
		}
		f.scripts[name] = &script{
			fs:       f.fs,
			io:       f.io,
			exec:     f.exec,
			srcPaths: srcPaths,
			name:     name,
			binPath:  filepath.Join(f.binDir, name+ext),
			rootDir:  f.rootDir,
		}
		f.names = append(f.names, name)
	}
	sort.Strings(f.names)
	return nil
}

func (f *scriptLoader) Get(name string) (script module.Script, ok bool) {
	script, ok = f.scripts[name]
	return
}

func (f *scriptLoader) Names() []string {
	return f.names
}

package cmd

import (
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/izumin5210/clig/pkg/clib"
	"github.com/x-izumin/grapi/pkg/grapicmd"
	"github.com/x-izumin/grapi/pkg/grapicmd/di"
	"github.com/x-izumin/grapi/pkg/grapicmd/internal/usecase"
)

func newInitCommand(ctx *grapicmd.Ctx) *cobra.Command {
	var cfg usecase.InitConfig

	cmd := &cobra.Command{
		Use:           "init [name]",
		Short:         "Initialize a grapi application",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root, err := parseInitArgs(ctx, args)
			if err != nil {
				return errors.WithStack(err)
			}
			zap.L().Debug("parseInitArgs", zap.String("root", root))

			u, err := di.NewInitializeProjectUsecase(ctx, clib.Path(root))
			if err != nil {
				return errors.WithStack(err)
			}

			return errors.WithStack(u.Perform(root, cfg))
		},
	}

	cmd.PersistentFlags().StringVar(&cfg.Revision, "revision", "", "Specify grapi revision")
	cmd.PersistentFlags().StringVar(&cfg.Branch, "branch", "", "Specify grapi branch")
	cmd.PersistentFlags().StringVar(&cfg.Version, "version", "", "Specify grapi version")
	cmd.PersistentFlags().BoolVar(&cfg.HEAD, "HEAD", false, "Use HEAD grapi")
	cmd.PersistentFlags().StringVar(&cfg.GrapiReplacementURL, "replace-grapi", "", "Specify grapi replacement url")
	cmd.PersistentFlags().StringVarP(&cfg.Package, "package", "p", "", `Package name of the application(default: "<parent_package_or_username>.<app_name>")`)
	cmd.PersistentFlags().BoolVar(&cfg.Dep, "use-dep", false, "Use dep instead of Modules")

	return cmd
}

func parseInitArgs(ctx *grapicmd.Ctx, args []string) (root string, err error) {
	if argCnt := len(args); argCnt != 1 {
		err = errors.Errorf("invalid argument count: want 1, got %d", argCnt)
		return
	}

	arg := args[0]
	root = ctx.RootDir.String()

	if arg == "." {
		return
	}
	root = arg
	if !filepath.IsAbs(arg) {
		root = ctx.RootDir.Join(arg).String()
	}
	return
}

package cmd

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/x-izumin/grapi/pkg/grapicmd"
	"github.com/x-izumin/grapi/pkg/grapicmd/di"
)

func newProtocCommand(ctx *grapicmd.Ctx) *cobra.Command {
	return &cobra.Command{
		Use:           "protoc",
		Short:         "Run protoc",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !ctx.IsInsideApp() {
				return errors.New("protoc command should be execute inside a grapi application directory")
			}
			protocw, err := di.NewProtocWrapper(ctx)
			if err != nil {
				return errors.WithStack(err)
			}
			return errors.WithStack(protocw.Exec(context.TODO()))
		},
	}
}

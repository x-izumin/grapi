package gencmd

import (
	newGencmd "github.com/x-izumin/grapi/pkg/gencmd"
)

// Option configures a command context.
type Option = newGencmd.Option

// WithGrapiCtx specifies a grapi command context.
var WithGrapiCtx = newGencmd.WithGrapiCtx

// WithCreateAppFunc specifies a dependencies initializer.
var WithCreateAppFunc = newGencmd.WithCreateAppFunc

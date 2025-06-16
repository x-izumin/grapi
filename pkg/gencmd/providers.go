package gencmd

import (
	"github.com/google/wire"
	newGencmd "github.com/x-izumin/grapi/pkg/gencmd"
)

var ProvideGrapiCtx = newGencmd.ProvideGrapiCtx
var ProvideCtx = newGencmd.ProvideCtx
var ProvideShouldRun = newGencmd.ProvideShouldRun
var ProvidePath = newGencmd.ProvidePath

// Set contains providers for DI.
var Set = wire.NewSet(newGencmd.Set)

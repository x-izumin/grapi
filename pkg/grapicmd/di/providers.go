package di

import (
	"github.com/google/wire"
	newDi "github.com/x-izumin/grapi/pkg/grapicmd/di"
)

var ProvideScriptLoader = newDi.ProvideScriptLoader

var ProvideGenerator = newDi.ProvideGenerator

var Set = wire.NewSet(newDi.Set)

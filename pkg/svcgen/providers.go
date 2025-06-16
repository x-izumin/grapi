package svcgen

import (
	"github.com/google/wire"
	newSvcgen "github.com/x-izumin/grapi/pkg/svcgen"
)

var ProvideParamsBuilder = newSvcgen.ProvideParamsBuilder

var Set = wire.NewSet(newSvcgen.Set)

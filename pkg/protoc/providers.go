package protoc

import (
	"github.com/google/wire"
	newProtoc "github.com/x-izumin/grapi/pkg/protoc"
)

var ProvideGexConfig = newProtoc.ProvideGexConfig

var ProvideToolRepository = newProtoc.ProvideToolRepository

// WrapperSet is a provider set that includes gex things and Wrapper instance.
var WrapperSet = wire.NewSet(newProtoc.WrapperSet)

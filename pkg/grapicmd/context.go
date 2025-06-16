package grapicmd

import (
	"github.com/google/wire"
	newGrapicmd "github.com/x-izumin/grapi/pkg/grapicmd"
)

// Ctx contains the runtime context of grpai.
type Ctx = newGrapicmd.Ctx

// Config stores general setting params and provides accessors for them.
type Config = newGrapicmd.Config

// CtxSet is a provider set that includes modules contained in Ctx.
var CtxSet = wire.NewSet(newGrapicmd.CtxSet)

var ProvideFS = newGrapicmd.ProvideFS
var ProvideViper = newGrapicmd.ProvideViper
var ProvideExec = newGrapicmd.ProvideExec
var ProvideIO = newGrapicmd.ProvideIO
var ProvideRootDir = newGrapicmd.ProvideRootDir
var ProvideConfig = newGrapicmd.ProvideConfig
var ProvideBuildConfig = newGrapicmd.ProvideBuildConfig
var ProvideProtocConfig = newGrapicmd.ProvideProtocConfig

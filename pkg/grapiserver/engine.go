package grapiserver

import (
	newGrapiserver "github.com/x-izumin/grapi/pkg/grapiserver"
)

// Engine is the framework instance.
type Engine = newGrapiserver.Engine

// New creates a server intstance.
var New = newGrapiserver.New

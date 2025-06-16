package grapiserver

import (
	newGrapiserver "github.com/x-izumin/grapi/pkg/grapiserver"
)

// Address represents a network end point address.
type Address = newGrapiserver.Address

type HTTPServerConfig = newGrapiserver.HTTPServerConfig

// Config contains configurations of gRPC and Gateway server.
type Config = newGrapiserver.Config

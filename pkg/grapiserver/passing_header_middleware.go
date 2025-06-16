package grapiserver

import (
	newGrapiserver "github.com/x-izumin/grapi/pkg/grapiserver"
)

// PassedHeaderDeciderFunc returns true if given header should be passed to gRPC server metadata.
type PassedHeaderDeciderFunc = newGrapiserver.PassedHeaderDeciderFunc

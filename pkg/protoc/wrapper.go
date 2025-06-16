package protoc

import (
	newProtoc "github.com/x-izumin/grapi/pkg/protoc"
)

// Wrapper can execute protoc commands for current project's proto files.
type Wrapper = newProtoc.Wrapper

// NewWrapper creates a new Wrapper instance.
var NewWrapper = newProtoc.NewWrapper

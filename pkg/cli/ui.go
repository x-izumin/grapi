package cli

import (
	newCli "github.com/x-izumin/grapi/pkg/cli"
)

// UI is an interface for intaracting with the terminal.
type UI = newCli.UI

// UIInstance retuens a singleton UI instance.
var UIInstance = newCli.UIInstance

// NewUI creates a new UI instance.
var NewUI = newCli.NewUI

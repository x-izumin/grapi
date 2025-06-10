package gencmd

import "github.com/x-izumin/grapi/pkg/cli"

// CreateAppFunc initializes dependencies.
type CreateAppFunc func(*Command) (*App, error)

// App contains dependencies to execute a generator.
type App struct {
	Generator Generator
	UI        cli.UI
}

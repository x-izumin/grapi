package gencmd

import (
	newGencmd "github.com/x-izumin/grapi/pkg/gencmd"
)

// Command represents a subcommand of a generator plugin. It will be converted to a *cobra.Command object internally.
type Command = newGencmd.Command

// File represents a file content.
type File = newGencmd.File

// Entry represents a file that will be generated.
type Entry = newGencmd.Entry

type ShouldRunFunc = newGencmd.ShouldRunFunc

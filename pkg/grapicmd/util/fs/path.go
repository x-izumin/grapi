package fs

import (
	newFs "github.com/x-izumin/grapi/pkg/grapicmd/util/fs"
)

const (
	// PackageSeparator is a package separator string on protobuf.
	PackageSeparator = newFs.PackageSeparator
)

// GetImportPath creates the golang package path from the given path.
var GetImportPath = newFs.GetImportPath

// GetPackageName generates the package name of this application from the given path and envs.
var GetPackageName = newFs.GetPackageName

// FindMainPackagesAndSources returns go source file names by main package directories.
var FindMainPackagesAndSources = newFs.FindMainPackagesAndSources

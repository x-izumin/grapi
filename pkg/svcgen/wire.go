//go:build wireinject
// +build wireinject

package svcgen

import (
	newSvcgen "github.com/x-izumin/grapi/pkg/svcgen"
)

var NewApp = newSvcgen.NewApp

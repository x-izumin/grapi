//go:build !wireinject
// +build !wireinject

package di

import (
	newDi "github.com/x-izumin/grapi/pkg/grapicmd/di"
)

var NewUI = newDi.NewUI

var NewScriptLoader = newDi.NewScriptLoader

var NewToolRepository = newDi.NewToolRepository

var NewProtocWrapper = newDi.NewProtocWrapper

var NewInitializeProjectUsecase = newDi.NewInitializeProjectUsecase

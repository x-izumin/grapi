//go:build wireinject
// +build wireinject

package di

var NewUI = newDi.NewUI

var NewScriptLoader = newDi.NewScriptLoader

var NewToolRepository = newDi.NewToolRepository

var NewProtocWrapper = newDi.NewProtocWrapper

var NewInitializeProjectUsecase = newDi.NewInitializeProjectUsecase

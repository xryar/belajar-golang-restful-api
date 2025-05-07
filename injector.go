//go:build wireinject
// +build wireinject

package main

import (
	"belajar-golang-retsful-api/app"
	"belajar-golang-retsful-api/controller"
	"belajar-golang-retsful-api/middleware"
	"belajar-golang-retsful-api/repository"
	"belajar-golang-retsful-api/service"
	myValidator "belajar-golang-retsful-api/validator"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		myValidator.NewValidator,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}

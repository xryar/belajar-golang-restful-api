package main

import (
	"belajar-golang-retsful-api/app"
	"belajar-golang-retsful-api/controller"
	"belajar-golang-retsful-api/helper"
	"belajar-golang-retsful-api/middleware"
	"belajar-golang-retsful-api/repository"
	"belajar-golang-retsful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

package main

import (
	"github.com/arvrao/proj-re-app/mod-re-app/controllers"
	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	app.Get("/", controllers.HelloWorld)
}

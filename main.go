package main

import (
	"os"

	"github.com/arvrao/proj-re-app/mod-re-app/initializers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	//err := godotenv.Load(".env")
	initializers.LoadEnvVariables()
}

func main() {
	app := fiber.New()
	AppRoutes(app)
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello world arvinda")
	// })
	println("Port:", os.Getenv("GO_PORT"))
	app.Listen(":" + os.Getenv("GO_PORT"))

}

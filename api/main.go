package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golangcimri/api/environ"
	"github.com/golangcimri/api/globals"
	"github.com/golangcimri/api/middlewares"
	"github.com/golangcimri/api/routes"
)

func init() {
	environ.Init("main")
}

func main() {
	app := fiber.New()

	middlewares.Fiber(app)

	routes.PrivateRoutes(app)

	err := app.Listen(fmt.Sprintf(":%s", globals.Variables.ApiPort))
	if err != nil {
		panic(err)
	}
}

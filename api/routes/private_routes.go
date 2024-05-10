package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golangcimri/api/controllers"
	"github.com/golangcimri/api/middlewares"
)

// Define the private routes
func PrivateRoutes(r *fiber.App) {

	route := r.Group("/api/v1/private")

	route.Get("/get/record/:id", middlewares.ApiTokenProtected, controllers.GetRecord)

}

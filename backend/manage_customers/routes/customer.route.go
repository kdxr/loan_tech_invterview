package routes

import (
	controller "manage_customers/controllers"
	"manage_customers/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitCustomerRoute(route fiber.Router) {

	route.Use(middlewares.Authenticate())

	route.Post("/create", controller.CreateCustomer)
	route.Get("/lists", controller.ListCustomers)
	route.Get("/info/:id", controller.InfoCustomer)
	route.Post("/update", controller.UpdateCustomer)
	route.Delete("/delete/:id", controller.DeleteCustomer)
}

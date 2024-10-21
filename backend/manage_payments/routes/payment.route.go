package routes

import (
	controller "manage_payments/controllers"
	"manage_payments/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitCustomerRoute(route fiber.Router) {

	route.Use(middlewares.Authenticate())

	route.Post("/create", controller.CreatePayment)
	route.Delete("/delete/:id", controller.DeletePayment)
}

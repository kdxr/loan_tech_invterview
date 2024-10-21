package routes

import (
	controller "manage_loan/controllers"
	"manage_loan/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitCustomerRoute(route fiber.Router) {

	route.Use(middlewares.Authenticate())

	route.Post("/create", controller.CreateLoan)
	route.Get("/lists", controller.ListLoans)
	route.Get("/info/:id", controller.InformationLoan)
	route.Post("/update", controller.UpdateLoan)
	route.Delete("/delete/:id", controller.DeleteLoan)
}

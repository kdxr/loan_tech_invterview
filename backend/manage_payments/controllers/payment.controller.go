package controller

import (
	"manage_payments/database"
	"manage_payments/models"
	"manage_payments/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreatePayment(c *fiber.Ctx) error {

	payload := &models.Payment{}

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Incorrect Parameter",
			"fields":  utils.ValidatorErrors(err),
		})
	}

	loan := &models.Loan{}

	if response := database.DB.First(&loan, payload.LoanId); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Loan Invalid",
		})
	}

	if result := database.DB.Create(&payload); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "success",
	})
}

func DeletePayment(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if response := database.DB.First(&models.Payment{}, id); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Payment Invalid",
		})
	}

	if response := database.DB.Delete(&models.Payment{}, id); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": response.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "success",
		"data":    "",
	})
}

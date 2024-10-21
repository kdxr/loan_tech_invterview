package controller

import (
	"manage_customers/database"
	"manage_customers/models"
	"manage_customers/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateCustomer(c *fiber.Ctx) error {

	payload := &models.Customer{}

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

func ListCustomers(c *fiber.Ctx) error {

	users := []models.Customer{}

	if response := database.DB.Select([]string{"id", "email", "name", "sex", "address", "telephonenumber", "tumbon", "district", "province", "status_active", "created_at"}).Find(&users); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": response.Error.Error(),
		})
	}

	// sanitizedUsers := []fiber.Map{}
	// for _, user := range users {
	// 	sanitizedUsers = append(sanitizedUsers, fiber.Map{
	// 		"email":           user.Email,
	// 		"telephoneNumber": user.TelphoneNumber,
	// 		"name":            user.Name,
	// 		"sex":             user.Sex,
	// 		"address":         user.Address,
	// 		"tumbon":          user.Tumbon,
	// 		"district":        user.District,
	// 		"province":        user.Province,
	// 		"statusActive":    user.StatusActive,
	// 		"createdAt":       user.CreatedAt,
	// 	})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "success",
		"data":    users,
	})
}

func InfoCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	customer := &models.Customer{}

	if response := database.DB.First(&customer, id); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": response.Error.Error(),
		})
	}

	result := fiber.Map{
		"email":           customer.Email,
		"telephoneNumber": customer.TelphoneNumber,
		"name":            customer.Name,
		"sex":             customer.Sex,
		"address":         customer.Address,
		"tumbon":          customer.Tumbon,
		"district":        customer.District,
		"province":        customer.Province,
		"statusActive":    customer.StatusActive,
		"createdAt":       customer.CreatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "success",
		"data":    result,
	})
}

func UpdateCustomer(c *fiber.Ctx) error {
	payload := &models.UpdateCustomer{}

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

	customer := &models.Customer{}

	if response := database.DB.First(&customer, payload.ID); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": response.Error.Error(),
		})
	}

	if response := database.DB.Model(&customer).Updates(&models.Customer{
		Email:          payload.Email,
		TelphoneNumber: payload.TelphoneNumber,
		Name:           payload.Name,
		Sex:            payload.Sex,
		Address:        payload.Address,
		Tumbon:         payload.Tumbon,
		District:       payload.District,
		Province:       payload.Province,
		StatusActive:   *payload.StatusActive,
	}); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": response.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "success",
	})
}

func DeleteCustomer(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if response := database.DB.Delete(&models.Customer{}, id); response.Error != nil {
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

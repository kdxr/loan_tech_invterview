package controller

import (
	"manage_loan/database"
	"manage_loan/models"
	"manage_loan/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateLoan(c *fiber.Ctx) error {

	payload := &models.Loan{}

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

	if response := database.DB.First(&customer, payload.CustomerId); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Customer Invalid",
		})
	}

	if result := database.DB.Model(&payload).Create(&payload); result.Error != nil {
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

func ListLoans(c *fiber.Ctx) error {
	loans := []models.LoanInformation{}

	if response := database.DB.Table("loans").Preload("Customer").Find(&loans); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": response.Error.Error(),
		})
	}

	results := []fiber.Map{}

	for _, loan := range loans {
		results = append(results, fiber.Map{
			"loanId":        loan.ID,
			"customerId":    loan.CustomerId,
			"customerEmail": loan.Customer.Email,
			"customerName":  loan.Customer.Name,
			"customerTel":   loan.Customer.TelphoneNumber,
			"loanAmount":    loan.LoanAmount,
			"interestRate":  loan.InterestRate,
			// "loanPayDate":   loan.PayDate.Format("02 January 2006 15:04:05"),
			// "loanStartDate": loan.StartDate.Format("02 January 2006 15:04:05"),
			// "loanEndDate":   loan.EndDate.Format("02 January 2006 15:04:05"),
			"loanStartDate": loan.StartDate,
			"loanEndDate":   loan.EndDate,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "success",
		"data":    results,
	})
}

func InformationLoan(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	loan := &models.LoanInformation{}

	if response := database.DB.Table("loans").Preload("Customer").First(&loan, id); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": response.Error.Error(),
		})
	}

	payments := &[]models.Payment{}

	if response := database.DB.Select("pay_amount", "principle_amount", "interest_amount", "created_at").Find(&payments, "loan_id = ?", id); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": response.Error.Error(),
		})
	}

	result := fiber.Map{
		"loanId":                 loan.ID,
		"customerId":             loan.CustomerId,
		"customerEmail":          loan.Customer.Email,
		"customerTelphoneNumber": loan.Customer.TelphoneNumber,
		"customerName":           loan.Customer.Name,
		"customerSex":            loan.Customer.Sex,
		"customerAddress":        loan.Customer.Address,
		"customerTumbon":         loan.Customer.Tumbon,
		"customerDistrict":       loan.Customer.District,
		"customerProvince":       loan.Customer.Province,
		"loanAmount":             loan.LoanAmount,
		"interestRate":           loan.InterestRate,
		// "loanPayDate":            loan.PayDate.Format("02 January 2006 15:04:05"),
		// "loanStartDate": loan.StartDate.Format("02 January 2006 15:04:05"),
		// "loanEndDate":   loan.EndDate.Format("02 January 2006 15:04:05"),
		"loanStartDate": loan.StartDate,
		"loanEndDate":   loan.EndDate,
		"payments":      payments,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "success",
		"data":    result,
	})
}

func UpdateLoan(c *fiber.Ctx) error {
	payload := &models.UpdateLoan{}

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

	if response := database.DB.First(&loan, payload.ID); response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": response.Error.Error(),
		})
	}

	if response := database.DB.Model(&loan).Updates(&models.Loan{
		LoanAmount:   payload.LoanAmount,
		InterestRate: payload.InterestRate,
		// PayDate:      payload.PayDate,
		StartDate: payload.StartDate,
		EndDate:   payload.EndDate,
	}); response.Error != nil {
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

func DeleteLoan(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if response := database.DB.Delete(&models.Loan{}, id); response.Error != nil {
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

package api

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type jsonResonse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var validate = validator.New()

func validateBody(c *fiber.Ctx, data any) error {

	println("Starting validation")

	if err := c.BodyParser(&data); err != nil {
		println("err1 in validation")
		return err
	}
	println("Data in validate body:", data)
	if err := validate.Struct(&data); err != nil {
		println("err2 in validation")
		return err
	}

	return nil
}

func writeJSON(c *fiber.Ctx, status int, data any) error {

	return c.JSON(jsonResonse{Status: status, Message: "Success", Data: data})
}

func errorJson(c *fiber.Ctx, err error, status ...int) error {

	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]

	}

	return c.JSON(jsonResonse{Status: statusCode, Message: err.Error()})
}

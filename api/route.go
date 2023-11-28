package api

import (
	"github.com/gofiber/fiber/v2"
)

func GetRoutes(app *fiber.App) {
	app.Post("/otp", sendSMS)
	app.Post("/verifyOTP", verifySMS)
}

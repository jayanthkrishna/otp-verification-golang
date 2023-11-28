package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jayanthkrishna/otp-verification-golang/data"
)

const apptimeout = time.Second * 5

func sendSMS(c *fiber.Ctx) error {

	_, cancel := context.WithTimeout(context.Background(), apptimeout)

	defer cancel()
	var payload data.OTPData

	// c.BindJSON(&payload)
	// validateBody(c, payload)

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	println("data in payload", payload.PhoneNumber)
	newData := data.OTPData{
		PhoneNumber: payload.PhoneNumber,
	}

	_, err := twillioSendOTP(newData.PhoneNumber)

	if err != nil {
		println("error :", err)
		return errorJson(c, err)

	}

	return writeJSON(c, http.StatusAccepted, "OTP Sent successfully")

}
func verifySMS(c *fiber.Ctx) error {

	_, cancel := context.WithTimeout(context.Background(), apptimeout)

	defer cancel()
	var payload data.VerifyData

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	fmt.Printf("Data in verify : %v\n", payload)
	newData := data.VerifyData{
		User: payload.User,
		Code: payload.Code,
	}

	err := twillioVerifyOTP(newData.User.PhoneNumber, newData.Code)

	if err != nil {
		return errorJson(c, err)

	}

	return writeJSON(c, http.StatusAccepted, "OTP verified successfully")

}

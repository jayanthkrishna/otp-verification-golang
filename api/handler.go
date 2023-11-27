package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jayanthkrishna/otp-verification-golang/data"
)

const apptimeout = time.Second * 5

func (app *Config) sendSMS() gin.HandlerFunc {

	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), apptimeout)

		defer cancel()
		var payload data.OTPData

		app.validateBody(c, payload)

		println("data in payload", payload.PhoneNumber)
		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twillioSendOTP(newData.PhoneNumber)

		if err != nil {
			app.errorJson(c, err)
			return

		}

		app.writeJSON(c, http.StatusAccepted, "OTP Sent successfully")

	}
}
func (app *Config) verifySMS() gin.HandlerFunc {

	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), apptimeout)

		defer cancel()
		var payload data.VerifyData

		app.validateBody(c, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		err := app.twillioVerifyOTP(newData.User.PhoneNumber, newData.Code)

		if err != nil {
			app.errorJson(c, err)
			return

		}

		app.writeJSON(c, http.StatusAccepted, "OTP verified successfully")

	}

}

package api

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envAccountsID(),
	Password: envAuthToken(),
})

func twillioSendOTP(phonenumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	println("Phone Number :", phonenumber)
	params.SetTo(phonenumber)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(envServiceID(), params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func twillioVerifyOTP(phonenumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}

	params.SetTo(phonenumber)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(envServiceID(), params)
	if err != nil {
		return err
	}

	if *resp.Status != "approved" {
		return errors.New("Not a valid code")
	}
	return nil

}

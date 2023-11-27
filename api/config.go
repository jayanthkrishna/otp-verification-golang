package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envAccountsID() string {
	println(godotenv.Unmarshal(".env"))

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func envServiceID() string {
	println(godotenv.Unmarshal(".env"))

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("TWILIO_SERVICE_ID")
}
func envAuthToken() string {
	println(godotenv.Unmarshal(".env"))

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("TWILIO_AUTHTOKEN")
}

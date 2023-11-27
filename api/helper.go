package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResonse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var validate = validator.New()

func (app *Config) validateBody(c *gin.Context, data any) error {

	println("Starting validation")

	if err := c.BindJSON(&data); err != nil {
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

func (app *Config) writeJSON(c *gin.Context, status int, data any) {

	c.JSON(status, jsonResonse{Status: status, Message: "Success", Data: data})
}

func (app *Config) errorJson(c *gin.Context, err error, status ...int) {

	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]

	}

	c.JSON(statusCode, jsonResonse{Status: statusCode, Message: err.Error()})
}

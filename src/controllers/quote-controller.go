package controllers

import (
	"github.com/InfoTCube/Seneca-the-Younger-API/src/entities"
	"github.com/InfoTCube/Seneca-the-Younger-API/src/services"
	"github.com/InfoTCube/Seneca-the-Younger-API/src/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type QuoteController interface {
	GetAll() []entities.Quote
	Add(ctx *gin.Context) error
}

type quoteController struct {
	service services.QuoteService
}

var validate *validator.Validate

func New(service services.QuoteService) QuoteController {
	validate = validator.New()
	validate.RegisterValidation("country-code", validators.ValidateCountryCode)
	return &quoteController{
		service: service,
	}
}

func (c *quoteController) GetAll() []entities.Quote {
	return c.service.GetAll()
}

func (c *quoteController) Add(ctx *gin.Context) error {
	var quote entities.Quote
	err := ctx.ShouldBindJSON(&quote)
	if err != nil {
		return err
	}
	err = validate.Struct(quote)
	if err != nil {
		return err
	}
	c.service.Add(quote)
	return nil
}

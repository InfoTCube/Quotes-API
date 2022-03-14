package controllers

import (
	"github.com/InfoTCube/Seneca-the-Younger-API/src/entity"
	"github.com/InfoTCube/Seneca-the-Younger-API/src/service"
	"github.com/gin-gonic/gin"
)

type QuoteController interface {
	GetAll() []entity.Quote
	Add(ctx *gin.Context) entity.Quote
}

type quoteController struct {
	service service.QuoteService
}

func New(service service.QuoteService) QuoteController {
	return &quoteController{
		service: service,
	}
}

func (c *quoteController) GetAll() []entity.Quote {
	return c.service.GetAll()
}

func (c *quoteController) Add(ctx *gin.Context) entity.Quote {
	var quote entity.Quote
	ctx.BindJSON(&quote)
	c.service.Add(quote)
	return quote
}

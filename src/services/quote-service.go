package services

import "github.com/InfoTCube/Seneca-the-Younger-API/src/entities"

type QuoteService interface {
	Add(entities.Quote) entities.Quote
	GetAll() []entities.Quote
	GetRandom() entities.Quote
	GetById(int32) entities.Quote
}

type quoteService struct {
	quotes []entities.Quote
}

func New() QuoteService {
	return &quoteService{}
}

func (service *quoteService) Add(quote entities.Quote) entities.Quote {
	service.quotes = append(service.quotes, quote)
	return quote
}

func (service *quoteService) GetAll() []entities.Quote {
	return service.quotes
}

func (service *quoteService) GetRandom() entities.Quote {
	return service.quotes[0]
}

func (service *quoteService) GetById(int32) entities.Quote {
	return service.quotes[0]
}

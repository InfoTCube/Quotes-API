package service

import "github.com/InfoTCube/Seneca-the-Younger-API/entity"

type QuoteService interface {
	Add(entity.Quote) entity.Quote
	GetAll() []entity.Quote
	GetRandom() entity.Quote
	GetById(int32) entity.Quote
}

type quoteService struct {
	quotes []entity.Quote
}

func New() QuoteService {
	return &quoteService{}
}

func (service *quoteService) Add(quote entity.Quote) entity.Quote {
	service.quotes = append(service.quotes, quote)
	return quote
}

func (service *quoteService) GetAll() []entity.Quote {
	return service.quotes
}

func (service *quoteService) GetRandom() entity.Quote {
	return service.quotes[0]
}

func (service *quoteService) GetById(int32) entity.Quote {
	return service.quotes[0]
}

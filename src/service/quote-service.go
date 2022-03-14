package service

import "github.com/InfoTCube/eneca-the-Younger-API/entity"

type QuoteService interface {
	Add(entity.Quote) entity.Quote
	GetAll() []entity.Quote
	GetRandom() entity.Quote
	GetById(int32) entity.Quote
}

type QuoteService struct {
	quotes []entity.Quote
}

func New() QuoteService {
	return &quoteService{}
}

func (service *quoteService) Add(entity.Quote) entity.Quote {

}

func (service *quoteService) GetAll() []entity.Quote {

}

func (service *quoteService) GetRandom() entity.Quote {

}

func (service *quoteService) GetById(int32) entity.Quote {

}

package dto

import (
	"go-scraper/entity"
	"go-scraper/helpers"
	"go-scraper/presentation"
)

func ScrapertoArray(payload []presentation.Product) presentation.Products {
	resp := presentation.Products{
		Products: payload,
	}

	return resp
}

func ProductRequest(payload presentation.Product) entity.Product {
	if payload.Description == "" {
		payload.Description = "-"
	}

	payload.Price = helpers.FormatCurrency(payload.Price)
	
	resp := entity.Product{
		Link:        payload.Link,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Rating:      payload.Rating,
		Merchant:    payload.Merchant,
	}

	return resp
}



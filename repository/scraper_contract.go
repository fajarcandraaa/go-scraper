package repository

import "go-scraper/entity"

type ScraperRepositoryContract interface {
	InsertData(payload entity.Product) error
}
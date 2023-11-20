package service

import "go-scraper/presentation"

type ScraperContract interface {
	MapScraping(start, end int) []presentation.Product
	ToCsv(payload []presentation.Product)
	StoreData(payload presentation.Products) error
}

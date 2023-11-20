package main

import (
	"go-scraper/config"
	"go-scraper/dto"
	"go-scraper/repository"
	"go-scraper/service"
	"log"

	"github.com/gocolly/colly/v2"
)

type Product struct {
	Link        string
	Name        string
	Description string
	Price       string
	Rating      string
	Merchant    string
}

func main() {
	var (
		startPage       = 1
		endPage         = 2
		c               = colly.NewCollector()
		dbconfiguration = config.Serve{}
	)

	db 			:= dbconfiguration.DBConfig()
	repo 		:= repository.NewScraperRepository(db)
	sClass 		:= service.NewScraper(c, repo)
	products 	:= sClass.MapScraping(startPage, endPage)
	payload 	:= dto.ScrapertoArray(products)
	err 		:= sClass.StoreData(payload)
	if err != nil {
		log.Fatalf("Error : %s", err.Error())
	}

	sClass.ToCsv(products)
}

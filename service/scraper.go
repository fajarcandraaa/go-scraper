package service

import (
	"encoding/csv"
	"fmt"
	"go-scraper/dto"
	"go-scraper/helpers"
	"go-scraper/presentation"
	"go-scraper/repository"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

type scraperService struct {
	col  *colly.Collector
	repo *repository.ScraperRepository
}

func NewScraper(col *colly.Collector, repo *repository.ScraperRepository) *scraperService {
	return &scraperService{
		col:  col,
		repo: repo,
	}
}

var _ ScraperContract = &scraperService{}

// Scraping implements ScraperContract.
func (s *scraperService) MapScraping(start, end int) []presentation.Product {
	var (
		wg        sync.WaitGroup
		products  = make([]presentation.Product, 0)
		item      = presentation.Product{}
		itemcount = 0
	)

	s.col.SetRequestTimeout(120 * time.Second)
	s.col.OnRequest(func(r *colly.Request) { helpers.SetAgent(r) })

	s.col.OnHTML("div.css-bk6tzz", func(parent *colly.HTMLElement) {
		wg.Add(5)
		go func() {
			parent.ForEach("a.css-54k5sq", func(i int, hrefScrap *colly.HTMLElement) {
				link := hrefScrap.Attr("href")
				item.Link = link

				if strings.HasPrefix(link, "https") {
					s.col.OnHTML("div[data-testid=lblPDPDescriptionProduk]", func(desc *colly.HTMLElement) {
						item.Description = desc.Text
					})

					s.col.OnRequest(func(r *colly.Request) {
						helpers.SetAgent(r)
					})

					err := s.col.Visit(link)
					if err != nil {
						log.Println("Error visiting link: ", link, err)
					}
				}
			})
			wg.Done()
		}()

		go func() {
			parent.ForEach("span.css-ywdpwd", func(i int, merName *colly.HTMLElement) {
				if i == 1 {
					nameMer := merName.Text
					item.Merchant = nameMer
				}
			})
			wg.Done()
		}()

		go func() {
			imgCount := 0
			parent.ForEach("div.css-1riykrk", func(i int, rateCount *colly.HTMLElement) {
				rateCount.ForEach("img.css-177n1u3", func(j int, rcount *colly.HTMLElement) {
					imgCount++
				})
			})
			item.Rating = strconv.Itoa(imgCount)
			wg.Done()
		}()

		go func() {
			parent.ForEach("div.css-11s9vse", func(i int, child1 *colly.HTMLElement) {
				item.Name = child1.Text
			})
			wg.Done()
		}()

		go func() {
			parent.ForEach("div.css-pp6b3e", func(i int, child2 *colly.HTMLElement) {
				item.Price = child2.ChildText("span")
			})
			wg.Done()
		}()

		wg.Wait()

		products = append(products, item)
		itemcount++
	})

	if itemcount == 101 {
		return products
	} else {
		for i := start; i <= end; i++ {
			err := s.col.Visit(fmt.Sprintf("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=5&page=%d", i))
			if err != nil {
				log.Fatal("MAIN LOG ERROR : ", err)
			}
		}

		fmt.Printf("\n TOTAL ITEM : %d \n", itemcount)
		return products
	}
}

// ToCsv implements ScraperContract.
func (*scraperService) ToCsv(payload []presentation.Product) {
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{
		"link",
		"name",
		"description",
		"price",
		"rating",
		"merchant",
	}
	writer.Write(headers)

	for _, prod := range payload {
		record := []string{
			prod.Link,
			prod.Name,
			prod.Description,
			prod.Price,
			prod.Rating,
			prod.Merchant,
		}

		writer.Write(record)
	}
	defer writer.Flush()
}

// StoreData implements ScraperContract.
func (s *scraperService) StoreData(payload presentation.Products) error {
	for _, r := range payload.Products {
		payload := dto.ProductRequest(r)
		err := s.repo.InsertData(payload)
		if err != nil {
			return err
		}
	}

	return nil
}

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

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
	c := colly.NewCollector()
	products := make([]Product, 0)

	// c.OnHTML("a.css-54k5sq", func(parent *colly.HTMLElement) {
	c.OnHTML("div.css-bk6tzz", func(parent *colly.HTMLElement) {
		imgCount := 0
		item := Product{}
		parent.ForEach("a.css-54k5sq", func(i int, hrefScrap *colly.HTMLElement) {
			link := hrefScrap.Attr("href")
			item.Link = link

			if strings.HasPrefix(link, "https") {
				c.OnHTML("div[data-testid=lblPDPDescriptionProduk]", func(desc *colly.HTMLElement) {
					item.Description = desc.Text
				})

				c.OnRequest(func(r *colly.Request) {
					r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
				})

				err := c.Visit(link)
				if err != nil {
					log.Println("Error visiting link: ", link, err)
				}
			}
		})
		parent.ForEach("span.css-ywdpwd", func(i int, merName *colly.HTMLElement) {
			if i == 1 {
				nameMer := merName.Text
				item.Merchant = nameMer
			}
		})

		parent.ForEach("div.css-1riykrk", func(i int, rateCount *colly.HTMLElement) {
			rateCount.ForEach("img.css-177n1u3", func(j int, rcount *colly.HTMLElement) {
				imgCount++
			})
		})

		item.Rating = strconv.Itoa(imgCount)

		parent.ForEach("div.css-11s9vse", func(i int, child1 *colly.HTMLElement) {
			item.Name = child1.Text

		})
		parent.ForEach("div.css-pp6b3e", func(i int, child2 *colly.HTMLElement) {
			item.Price = child2.ChildText("span")
		})

		products = append(products, item)
	})

	c.SetRequestTimeout(120 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		js, err := json.MarshalIndent(products, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("products.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}

	})

	for i := 1; i <= 2; i++ {
		err := c.Visit(fmt.Sprintf("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=5&page=%d", i))
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// defining the CSV headers
	headers := []string{
		"link",
		"name",
		"description",
		"price",
		"rating",
		"merchant",
	}
	// writing the column headers
	writer.Write(headers)

	// adding each Pokemon product to the CSV output file
	for _, prod := range products {
		// converting a PokemonProduct to an array of strings
		record := []string{
			prod.Link,
			prod.Name,
			prod.Description,
			prod.Price,
			prod.Rating,
			prod.Merchant,
		}

		// writing a new CSV record
		writer.Write(record)
	}
	defer writer.Flush()
}

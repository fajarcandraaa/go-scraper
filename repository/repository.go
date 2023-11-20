package repository

import "github.com/jinzhu/gorm"

type Repository struct{
	ScraperRepo ScraperRepositoryContract
}

func NewScraper(db *gorm.DB) ScraperRepositoryContract {
	return NewScraperRepository(db)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{}
}

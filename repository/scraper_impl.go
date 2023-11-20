package repository

import (
	"go-scraper/entity"

	"github.com/jinzhu/gorm"
)

type ScraperRepository struct {
	db *gorm.DB
}

func NewScraperRepository(db *gorm.DB) *ScraperRepository {
	return &ScraperRepository{
		db: db,
	}
}

var _ ScraperRepositoryContract = &ScraperRepository{}

// InsertData implements ScraperRepositoryContract.
func (r *ScraperRepository) InsertData(payload entity.Product) error {
	var (
		query = `
			INSERT INTO products (link,name,description,price,rating,merchant)
			VALUES ($1, $2, $3, $4, $5, $6)
		`
	)

	arg := []interface{}{
		&payload.Link,
		&payload.Name,
		&payload.Description,
		&payload.Price,
		&payload.Rating,
		&payload.Merchant,
	}

	err := r.db.Exec(query, arg...).Error
	if err != nil {
		return err
	}

	return nil
}

// Package db provides db
package db

import (
	"time"

	models "github.com/abodsakah/BTH-appen/backend/api/src/Models"
	"gorm.io/gorm"
)

// CreateNews function
func CreateNews(db *gorm.DB, news *models.News) error {
	// set creation date
	news.CreatedAt = time.Now()

	// create news in database
	err := db.Create(&news).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteNews function
// Takes a news ID and deletes the news from the database.
//
// Or returns an error.
func DeleteNews(db *gorm.DB, newsID uint) (models.News, error) {
	// find news
	news := models.News{}
	err := db.Where("id = ?", newsID).First(&news).Error
	if err != nil {
		return models.News{}, err
	}
	// delete news from database
	err = db.Delete(&news).Error
	if err != nil {
		return models.News{}, err
	}
	return news, nil
}

// GetNews function
// Returns an array with all news
//
// Or returns an error.
func GetNews(db *gorm.DB) ([]models.News, error) {
	var news []models.News
	err := db.Order("id ASC").Find(&news).Error
	if err != nil {
		return nil, err
	}
	return news, nil
}

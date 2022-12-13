// Package db provides db
package db

import (
	"time"

	"gorm.io/gorm"
)

// CreateNews function
func CreateNews(db *gorm.DB, news *News) error {
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
func DeleteNews(db *gorm.DB, newsID uint) (News, error) {
	// find news
	news := News{}
	err := db.Where("id = ?", newsID).First(&news).Error
	if err != nil {
		return News{}, err
	}
	// delete news from database
	err = db.Delete(&news).Error
	if err != nil {
		return News{}, err
	}
	return news, nil
}

// GetNews function
// Returns an array with all news
//
// Or returns an error.
func GetNews(db *gorm.DB) ([]News, error) {
	var news []News
	err := db.Order("id ASC").Find(&news).Error
	if err != nil {
		return nil, err
	}
	return news, nil
}

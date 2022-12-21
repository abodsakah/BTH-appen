// Package scraper provides scraper
package scraper

import (
	"log"
	"strings"
	"time"

	"github.com/abodsakah/BTH-appen/backend/api/src/DB"
	"github.com/abodsakah/BTH-appen/backend/api/src/Notifications"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
)

// domain and newsURL are the domain and the url to the news
const (
	domain  = "www.bth.se"
	newsURL = "https://www.bth.se/category/nyheter"
)

// Start function
// Runs GetNews() every 5 hours.
func Start(gormDB *gorm.DB) {
	for {
		// slice to keep track of new articles
		var newArticles []db.News
		// scrape all articles
		news, err := GetNews()
		if err != nil {
			log.Println("scraper; Something went wrong scraping data. trying again in 1 hour ")
			time.Sleep(time.Hour)
			continue
		}
		// Add the news article to the database
		for _, article := range news {
			err := db.CreateNews(gormDB, &article)
			if err != nil {
				continue
			}
			newArticles = append(newArticles, article)
		}
		// send notification to users for each new article that was added to the database.
		err = notifications.SendNewsPushMessage(gormDB, newArticles)
		if err != nil {
			log.Println(err)
		}

		time.Sleep(5 * time.Hour)
	}
}

// GetNews function to get news from the website
func GetNews() ([]db.News, error) {
	// articles slice
	var news []db.News

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	// setup scraper logic
	c.OnHTML(".Article-result", func(h *colly.HTMLElement) {
		// Get the title
		h.ForEach(".ArticleItem", func(_ int, e *colly.HTMLElement) {
			var err error
			article := db.News{}
			article.Title = e.ChildText("h2")

			// Get the date
			e.ForEach("p", func(_ int, e *colly.HTMLElement) {
				if strings.Contains(e.Text, "Publicerad") {
					date := e.Text
					date = strings.Split(date, " ")[1]
					// format as: yyyy-mm-dd
					article.Date, err = time.Parse("2006-01-02", date)
					if err != nil {
						log.Println(err)
					}
				}
			})
			// Get the description
			e.ForEach("div[class=article-category-page]", func(_ int, description *colly.HTMLElement) {
				selction := description.DOM
				childNodes := selction.Children().Nodes
				if len(childNodes) > 1 {
					article.Description = selction.FindNodes(childNodes[1]).Text()
				} else if len(childNodes) == 1 {
					article.Description = selction.FindNodes(childNodes[0]).Text()
				}
			})
			// Get the link
			article.Link = e.ChildAttr("a", "href")
			news = append(news, article)
		})
	})

	// Start scraping on https://www.bth.se/category/nyheter
	err := c.Visit(newsURL)
	if err != nil {
		return nil, err
	}
	return news, nil
}

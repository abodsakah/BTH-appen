package scraper

import (
	"fmt"
	"log"
	"strings"
	"time"

	db "github.com/abodsakah/BTH-appen/backend/api/src/DB"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
)

// Start function to get the script sleep for 5 hours
func Start(gormDB *gorm.DB) {
	for {
		GetNews(gormDB)
		time.Sleep(5 * time.Hour)
	}
}

// GetNews function to get news from the website
func GetNews(gormDB *gorm.DB) {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.bth.se"),
	)

	c.OnHTML(".Article-result", func(h *colly.HTMLElement) {
		// Get the title
		h.ForEach(".ArticleItem", func(_ int, e *colly.HTMLElement) {
			title := e.ChildText("h2")
			fmt.Println(title)
			var date string
			var des string

			// Get the date
			e.ForEach("p", func(_ int, e *colly.HTMLElement) {
				if strings.Contains(e.Text, "Publicerad") {
					date = e.Text
					fmt.Println(date)
				}
			})
			// Get the description
			e.ForEach("div[class=article-category-page]", func(_ int, description *colly.HTMLElement) {
				selction := description.DOM
				childNodes := selction.Children().Nodes
				if len(childNodes) > 1 {
					des = selction.FindNodes(childNodes[1]).Text()
				} else if len(childNodes) == 1 {
					des = selction.FindNodes(childNodes[0]).Text()
				}
			})
			// Get the link
			link := e.ChildAttr("a", "href")
			fmt.Println(link)

			article := db.News{
				Title:       title,
				Date:        date,
				Description: des,
				Link:        link,
			}

			err := db.CreateNews(gormDB, &article) // Create the news in the database
			if err != nil {
				log.Println(err)
			}
		})
	})

	err := c.Visit("https://www.bth.se/category/nyheter")
	if err != nil {
		log.Println(err)
	}
}

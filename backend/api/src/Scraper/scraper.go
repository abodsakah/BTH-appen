package scraper

import (
	"encoding/json"
	"fmt"
	"strings"

	db "github.com/abodsakah/BTH-appen/backend/api/src/DB"
	"github.com/gocolly/colly"
)

func GetNews() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.bth.se"),
	)

	var newsList []db.News

	c.OnHTML(".Article-result", func(h *colly.HTMLElement) {

		h.ForEach(".ArticleItem", func(_ int, e *colly.HTMLElement) {
			title := e.ChildText("h2")
			fmt.Println(title)
			var date string
			var des string

			e.ForEach("p", func(_ int, e *colly.HTMLElement) {
				if strings.Contains(e.Text, "Publicerad") {
					date = e.Text
					fmt.Println(date)
				}
			})
			e.ForEach("div[class=article-category-page]", func(_ int, description *colly.HTMLElement) {
				selction := description.DOM
				childNodes := selction.Children().Nodes
				if len(childNodes) > 1 {
					des = selction.FindNodes(childNodes[1]).Text()
				} else if len(childNodes) == 1 {
					des = selction.FindNodes(childNodes[0]).Text()
				}
			})

			article := db.News{
				Title:       title,
				Date:        date,
				Description: des,
				Link:        "Bth.se",
			}

			newsList = append(newsList, article)
		})

	})

	c.Visit("https://www.bth.se/category/nyheter")
	results, _ := json.MarshalIndent(&newsList, "", "\n")

	fmt.Printf("%s", results)

}

/*
	// Get title
	title := h.ChildText("h2")
	newsList = append(newsList, db.News{Title: title})

	// Get date
	h.ForEach("p", func(_ int, d *colly.HTMLElement) {
		if strings.Contains(d.Text, "Publicerad") {
			date := d.Text
			newsList = append(newsList, db.News{Date: date})
		}
	})

	// Get description
	h.ForEach("div[class=article-category-page]", func(_ int, description *colly.HTMLElement) {
		selction := description.DOM
		childNodes := selction.Children().Nodes
		if len(childNodes) > 1 {
			des := selction.FindNodes(childNodes[1]).Text()
			newsList = append(newsList, db.News{Description: des})
		} else if len(childNodes) == 1 {
			des := selction.FindNodes(childNodes[0]).Text()
			newsList = append(newsList, db.News{Description: des})
		}
	})

	// Get link
	h.ForEach()
	link := h.Attr("href")
	newsList = append(newsList, db.News{Link: link})
*/

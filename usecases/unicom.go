package usecases

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"peach-core/entities"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapeUnicom() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("www.unicom.com.uy", "unicom.com.uy", "https://www.unicom.com.uy/"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./unicom_cache"),
	)

	//FIN LOGIN
	/*err := c.Post("https://www.unicom.com.uy/Account/Login", map[string]string{"Username": "user", "Password": "pass"})
	if err != nil {
		log.Fatal(err)
	}*/

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Create another collector to scrape course details
	detailCollector := c.Clone()

	products := make([]entities.Product, 0, 200)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// If attribute class is this long string return from callback
		// As this a is irrelevant
		if e.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
			return
		}
		link := e.Attr("href")
		// If link start with browse or includes either signup or login return from callback
		if !strings.HasPrefix(link, "/Producto?id=") || strings.Index(link, "=signup") > -1 || strings.Index(link, "=login") > -1 {
			return
		}
		// start scaping the page under the link found
		e.Request.Visit(link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	// On every a HTML element which has name attribute call callback
	c.OnHTML(`a[name]`, func(e *colly.HTMLElement) {
		// Activate detailCollector if the link contains "coursera.org/learn"
		productURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(productURL, "https://www.unicom.com.uy/Producto?id=") != -1 {
			detailCollector.Visit(productURL)
		}
	})

	// Extract details of the course
	detailCollector.OnHTML(`div[id=rendered-content]`, func(e *colly.HTMLElement) {
		log.Println("Product found", e.Request.URL)
		product := entities.Product{
			ProductID: "1",
			Provider:  1,
			Link:      e.Request.URL.String(),
			Brand:     "emptyfornow",
			Name:      "emptyfornow",
			Stock:     false,
			Category:  "emptyfornow",
			Price:     123,
		}
		// fill products with real info, then store in db
		// Iterate over rows of the table which contains different information
		// about the course
		products = append(products, product)
	})

	// Start scraping on http://coursera.com/browse
	c.Visit("https://www.unicom.com.uy/")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(products)
}

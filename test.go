package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Function to search DuckDuckGo and scrape multiple search results
func searchDuckDuckGo(query string) ([]string, error) {
	baseURL := "https://duckduckgo.com/html/"

	// Encode the search query
	data := url.Values{}
	data.Set("q", query)

	// Send the HTTP POST request (DuckDuckGo search form uses POST)
	resp, err := http.PostForm(baseURL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the HTML response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Find search result links and titles
	var results []string
	doc.Find(".result__title a").Each(func(i int, s *goquery.Selection) {
		// Extract the title text and the URL
		title := strings.TrimSpace(s.Text())
		link, exists := s.Attr("href")
		if exists {
			results = append(results, fmt.Sprintf("%d. %s (%s)", i+1, title, link))
		}
	})

	return results, nil
}

func main() {
	// Example search query
	query := "golang tutorials"

	// Perform the search
	results, err := searchDuckDuckGo(query)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Display the search results
	fmt.Printf("Search Results for '%s':\n", query)
	for _, result := range results {
		fmt.Println(result)
	}

	if len(results) == 0 {
		fmt.Println("No results found.")
	}
}

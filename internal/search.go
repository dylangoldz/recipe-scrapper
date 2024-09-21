package internal

import (
	"fmt"
	"os"

	g "github.com/serpapi/google-search-results-golang"
)

func Search(query []string) {
	apiKey := os.Getenv("SERP_API_KEY")
	for _, q := range query {
		parameter := map[string]string{
			"q":      q,
			"engine": "google",
			"num":    "1",
		}
		search := g.NewGoogleSearch(parameter, apiKey)
		data, err := search.GetJSON()
		results := data["organic_results"].([]interface{})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(results)

	}
}

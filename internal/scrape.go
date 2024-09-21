package internal

import (
	"fmt"
	"time"

	scrapper "github.com/kkyr/go-recipe/pkg/recipe"
)

type recipe struct {
	Name         string
	CookTime     time.Duration
	Ingredients  []string
	Instructions []string
	ImageURL     string
	Author       string
}

// potentially used for custom scrapping if serp api is not enough
func Scrape(urls []string) {
	for _, url := range urls {
		r, err := scrapper.ScrapeURL(url)
		if err != nil {
			fmt.Println(err)
		}

		recipeData := recipe{}

		recipeData.Ingredients, _ = r.Ingredients()
		recipeData.Instructions, _ = r.Instructions()
		recipeData.Author, _ = r.Author()
		recipeData.Name, _ = r.Name()
		recipeData.ImageURL, _ = r.ImageURL()
		recipeData.CookTime, _ = r.CookTime()

		fmt.Printf("Name: %s\n\nCook Time: %s\n\nIngredients: %v\n\nInstructions: %v\n\nImage URL: %s\n\nAuthor: %s\n\n",
			recipeData.Name, recipeData.CookTime,
			recipeData.Ingredients, recipeData.Instructions,
			recipeData.ImageURL, recipeData.Author)
	}
}

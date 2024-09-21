package main

import (
	"fmt"

	"github.com/dylangoldz/recipe-scrapper/internal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	queries := []string{
		"Italian cuisine",
		"Mexican cuisine",
		"Chinese cuisine",
		"Indian cuisine",
		"French cuisine",
		"Japanese cuisine",
		"Thai cuisine",
		"Greek cuisine",
		"Mediterranean cuisine",
		"American cuisine",
		"chicken recipes",
		"pasta dishes",
		"fish recipes",
		"vegetarian meals",
		"beef dishes",
		"pork recipes",
		"seafood dishes",
		"salad ideas",
		"soup recipes",
		"dessert recipes",
	}
	internal.Search(queries)
}

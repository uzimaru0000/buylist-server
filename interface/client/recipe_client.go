package client

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type recipeClient struct {
}

type RecipeClient interface {
	Get(url string) ([]string, error)
}

func NewClient() RecipeClient {
	return &recipeClient{}
}

func (client *recipeClient) Get(url string) ([]string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return []string{}, err
	}

	strs := doc.Find("div.ingredient").Map(func(i int, s *goquery.Selection) string {
		nameNode := s.Find("div.ingredient_name > span.name")
		name := strings.Replace(nameNode.Text(), "\n", "", 1)
		if name == "" {
			return ""
		}

		return name
	})

	return strs, nil

}

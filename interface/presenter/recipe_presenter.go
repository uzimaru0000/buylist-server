package presenter

import (
	"log"
	"strconv"
	"sync"

	"github.com/uzimaru0000/buylist/interface/client"
	"github.com/uzimaru0000/buylist/interface/perser"
	"github.com/uzimaru0000/buylist/usecase/presenter"
)

type recipePresenter struct {
	client client.RecipeClient
	perser perser.Perser
}

type RecipePresenter interface {
	presenter.RecipePresenter
}

func NewRecipePresenter(client client.RecipeClient, perser perser.Perser) RecipePresenter {
	return &recipePresenter{client, perser}
}

func (presenter *recipePresenter) Responce(urls []string) (*map[string]string, error) {
	strs, err := presenter.multiRequest(urls)
	if err != nil {
		return nil, err
	}
	data := presenter.perser.Perse(strs)

	ingre := make(map[string]string)
	for key, val := range data {
		ingre[key] = strconv.FormatInt(int64(val.Amount), 10) + val.Unit
	}

	return &ingre, nil
}

func (presenter *recipePresenter) multiRequest(urls []string) ([]string, error) {
	values := make(chan []string)
	finish := make(chan bool)
	isError := make(chan error)
	result := make(chan []string)

	go func() {
		val := make([]string, 0)

		for {
			select {
			case v := <-values:
				val = append(val, v...)
			case <-isError:
				return
			case <-finish:
				result <- val
				return
			}
		}
	}()

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for _, url := range urls {
			go func(str string) {
				log.Printf("fetch start")
				strs, err := presenter.client.Get(str)
				if err != nil {
					isError <- err
				}
				values <- strs
				log.Printf("fetch end")
				wg.Done()
			}(url)
		}
	}()

	wg.Wait()
	finish <- true

	err := <-isError
	res := <-result
	if err != nil {
		return nil, err
	}
	return res, nil

}

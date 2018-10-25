package presenter

import (
	"sync"

	"github.com/uzimaru0000/buylist/interface/client"
	"github.com/uzimaru0000/buylist/usecase/presenter"
)

type recipePresenter struct {
	client client.RecipeClient
}

type RecipePresenter interface {
	presenter.RecipePresenter
}

func NewRecipePresenter(client client.RecipeClient) RecipePresenter {
	return &recipePresenter{client}
}

func (presenter *recipePresenter) Responce(urls []string) ([]string, error) {
	strs, err := presenter.multiRequest(urls)
	if err != nil {
		return nil, err
	}

	response := make([]string, 0)

	for _, val := range strs {
		if val != "" {
			response = append(response, val)
		}
	}

	return response, nil
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
		for _, url := range urls {
			wg.Add(1)
			go func(str string) {
				strs, err := presenter.client.Get(str)
				if err != nil {
					isError <- err
				} else {
					values <- strs
				}
				wg.Done()
			}(url)
		}
		wg.Wait()
		finish <- true
	}()

	select {
	case r := <-result:
		return r, nil
	case err := <-isError:
		return nil, err
	}

}

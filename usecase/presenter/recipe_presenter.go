package presenter

type RecipePresenter interface {
	Responce(urls []string) ([]string, error)
}

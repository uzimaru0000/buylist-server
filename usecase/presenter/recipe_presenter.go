package presenter

type RecipePresenter interface {
	Response(urls []string) ([]string, error)
}

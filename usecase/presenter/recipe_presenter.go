package presenter

type RecipePresenter interface {
	Responce(urls []string) (*map[string]string, error)
}

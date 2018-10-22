package storage

import (
	"cloud.google.com/go/firestore"
	"github.com/uzimaru0000/buylist/domain/model"
	"github.com/uzimaru0000/buylist/usecase/repository"
	"golang.org/x/net/context"
)

type storage struct {
	firestore *firestore.Client
}

type Storage interface {
	repository.BuyListRepository
}

func NewStorage(firestore *firestore.Client) Storage {
	return &storage{firestore}
}

func (storage *storage) Store(list *model.BuyList) (*model.BuyList, error) {
	ref, _, err := storage.firestore.Collection("buylist").Add(context.Background(), list.Ingredients)

	list.ID = ref.ID
	return list, err
}

func (storage *storage) Find(list *model.BuyList) (*model.BuyList, error) {
	ref := storage.firestore.Collection("buylist").Doc(list.ID)
	ss, err := ref.Get(context.Background())
	if err != nil {
		return nil, err
	}

	list.Ingredients = convert(ss.Data())
	return list, nil
}

func convert(data map[string]interface{}) map[string]string {
	result := make(map[string]string)

	for key, val := range data {
		result[key] = val.(string)
	}

	return result
}

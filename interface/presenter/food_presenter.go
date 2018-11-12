package presenter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/uzimaru0000/buylist/domain/model"
	"github.com/uzimaru0000/buylist/usecase/presenter"
)

type foodPresenter struct {
	AppKey string
}

type FoodPresenter interface {
	presenter.FoodPresenter
}

func NewFoodPresenter(appkey string) FoodPresenter {
	return &foodPresenter{AppKey: appkey}
}

func (p *foodPresenter) Get(code int) (*model.Food, error) {
	values := url.Values{}
	values.Add("appid", p.AppKey)
	values.Add("jan", strconv.Itoa(code))

	url := "http://shopping.yahooapis.jp/ShoppingWebService/V1/json/itemSearch"

	res, err := http.Get(url + "?" + values.Encode())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data struct {
		ResultSet struct {
			Zero struct {
				Result struct {
					Zero struct {
						Name  string
						Image map[string]string
					} `json:"0"`
				}
			} `json:"0"`
		}
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	food := &model.Food{}

	food.Name = data.ResultSet.Zero.Result.Zero.Name
	food.ImageURL = data.ResultSet.Zero.Result.Zero.Image["Medium"]
	food.Exp = time.Now().UnixNano() / int64(time.Millisecond)
	food.Amount = 1

	return food, nil
}

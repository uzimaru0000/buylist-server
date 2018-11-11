package presenter_test

import (
	"testing"

	"github.com/uzimaru0000/buylist/domain/model"

	"github.com/uzimaru0000/buylist/interface/presenter"
)

const ApiKey = "YOUR API KEY"

func TestGetFood(t *testing.T) {
	pre := presenter.NewFoodPresenter(ApiKey)

	food, err := pre.Get(4901330575557)
	if err != nil {
		t.Fatal(err)
	}

	ans := &model.Food{Name: "カルビー じゃがりこ 塩とごま油味 52g×12個", ImageURL: "https://item-shopping.c.yimg.jp/i/g/dansyakudou_das01m710fqr"}

	if *food != *ans {
		t.Fatalf("Error : %v\n", *food)
	}

	t.Logf("%v\n", *food)

}

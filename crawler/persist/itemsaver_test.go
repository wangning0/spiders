package persist

import (
	"testing"
	"github.com/wangning0/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
)

func TestSave(t *testing.T) {
	expecetd := model.Profile{
		Name:       "阿兰",
		Gender:     "女",
		Age:        27,
		Height:     158,
		Weight:     0,
		Income:     "3001-5000元",
		Marriage:   "未婚",
		Education:  "中专",
		Occupation: "--",
		Hukou:      "四川阿坝",
		Xinzuo:    "双子座",
		House:      "租房",
		Car:        "未购车",
	}
	id, err := save(expecetd)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	resp, err := client.Get().Index("dataing_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	var actual model.Profile
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	if actual != expecetd {
		t.Errorf("check error")
	}
}
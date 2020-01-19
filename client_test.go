package gostic_test

import (
	"errors"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/toyaha/gostic"
	"strings"
	"testing"
)

func TestClient_Do(t *testing.T) {
	t.Run("search", func(t *testing.T) {
		es, err := gostic.NewClientDefault()
		if err != nil {
			t.Error(err)
		}
		str := `{"query": {"match_all": {}}}`
		req := esapi.SearchRequest{
			Index: []string{"idx"},
			Body:  strings.NewReader(str),
		}
		res, err := es.Do(req)
		if err != nil {
			t.Error(err)
		}

		var target struct{
			Hits struct {
				Hits []struct {
					Source struct{
						Id int `json:"id"`
						Name string `json:"name"`
					} `json:"_source"`
				} `json:"hits"`
			} `json:"hits"`
		}
		err = res.GetBodyStruct(&target)
		if err != nil {
			t.Error(err)
		}
		if len(target.Hits.Hits) < 1 {
			t.Error(errors.New("length"))
			return
		}

		{
			check := 1
			if target.Hits.Hits[0].Source.Id != check {
				t.Error("target:", target.Hits.Hits[0].Source.Id)
				t.Error("check :", check)
				return
			}
		}
		{
			check := "abc"
			if target.Hits.Hits[0].Source.Name != check {
				t.Error("target:", target.Hits.Hits[0].Source.Name)
				t.Error("check :", check)
				return
			}
		}
	})
}

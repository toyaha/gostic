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
			return
		}
		str := `{"query": {"match_all": {}}}`
		req := esapi.SearchRequest{
			Index: []string{"idx"},
			Body:  strings.NewReader(str),
		}
		res, err := es.Do(req)
		if err != nil {
			t.Error(err)
			return
		}

		var target struct {
			Hits struct {
				Hits []struct {
					Source struct {
						Id   int    `json:"id"`
						Name string `json:"name"`
					} `json:"_source"`
				} `json:"hits"`
			} `json:"hits"`
		}
		err = res.GetBodyStruct(&target)
		if err != nil {
			t.Error(err)
			return
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

func TestClient_DoWithGetBodyMap(t *testing.T) {
	t.Run("aggregations", func(t *testing.T) {
		var err error

		idx := "idx"
		es, err := gostic.NewClientDefault()
		if err != nil {
			t.Error(err)
			return
		}

		str := `{"size":0,"aggs":{"group":{"max":{"field":"id"}}}}`
		req := esapi.SearchRequest{
			Index: []string{idx},
			Body:  strings.NewReader(str),
		}
		result, _, err := es.DoWithGetBodyMap(req)
		if err != nil {
			t.Error(err)
			return
		}

		{
			a := result["aggregations"].(map[string]interface{})
			b := a["group"].(map[string]interface{})
			c := b["value"].(float64)
			target := c
			check := 2.0
			if target != check {
				t.Error("target:", target)
				t.Error("check :", check)
				return
			}
		}
	})
}

func TestClient_DoWithGetBodyStruct(t *testing.T) {
	t.Run("aggregations", func(t *testing.T) {
		var err error

		idx := "idx"
		es, err := gostic.NewClientDefault()
		if err != nil {
			t.Error(err)
			return
		}

		str := `{"size":0,"aggs":{"group":{"max":{"field":"id"}}}}`
		req := esapi.SearchRequest{
			Index: []string{idx},
			Body:  strings.NewReader(str),
		}
		var result struct {
			Aggregations struct {
				Group struct {
					Value float64 `json:"value"`
				} `json:"group"`
			} `json:"aggregations"`
		}
		_, err = es.DoWithGetBodyStruct(req, &result)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := result.Aggregations.Group.Value
			check := 2.0
			if target != check {
				t.Error("target:", target)
				t.Error("check :", check)
				return
			}
		}
	})
}

package gostic

import (
	"context"
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
)

type Client struct {
	Client *elasticsearch.Client
	Config *Config
}

func (rec *Client) Init(config *Config) error {
	if config == nil {
		config = &Config{}
		config.Init(nil)
	}
	rec.Config = config

	{
		client, err := elasticsearch.NewClient(*rec.Config.ElasticConfig)
		if err != nil {
			return err
		}
		rec.Client = client
	}

	return nil
}

func (rec *Client) Do(request esapi.Request) (*Response, error) {
	res, err := request.Do(context.Background(), rec.Client.Transport)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		str, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.New("response error")
		}
		return nil, errors.New(string(str))
	}
	defer func() {
		_ = res.Body.Close()
	}()

	response := &Response{}
	err = response.Init(res)
	if err != nil {
		return nil, err
	}

	bodyMap, err := response.GetBodyMap()
	if err != nil {
		return nil, err
	}

	if val, ok := bodyMap["errors"]; ok {
		if val.(bool) {
			return nil, errors.New(response.GetBodyString())
		}
	}

	return response, nil
}

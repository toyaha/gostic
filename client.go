package gostic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
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
		by := &bytes.Buffer{}
		_, err := io.Copy(by, res.Body)
		if err != nil {
			return nil, errors.New("response error")
		}
		return nil, errors.New(string(by.Bytes()))
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

func (rec *Client) DoWithGetBody(request esapi.Request) ([]byte, error) {
	response, err := rec.Do(request)
	if err != nil {
		return nil, err
	}

	by := response.GetBody()

	return by, nil
}

func (rec *Client) DoWithGetBodyMap(request esapi.Request) (map[string]interface{}, error) {
	by, err := rec.DoWithGetBody(request)
	if err != nil {
		return nil, err
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(by, &responseMap)
	if err != nil {
		return nil, err
	}

	return responseMap, nil
}

func (rec *Client) DoWithGetBodyString(request esapi.Request) (string, error) {
	by, err := rec.DoWithGetBody(request)
	if err != nil {
		return "", err
	}

	return string(by), nil
}

func (rec *Client) DoWithGetBodyStruct(request esapi.Request, dest interface{}) error {
	by, err := rec.DoWithGetBody(request)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(by)
	err = json.NewDecoder(reader).Decode(&dest)
	if err != nil {
		return err
	}

	return nil
}

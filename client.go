package gostic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func NewClient(config *Config) (*Client, error) {
	var client = &Client{}

	err := func() error {
		if config == nil {
			config = NewConfig(nil)
		}
		client.Config = config

		if client.Config.ElasticConfig == nil {
			return errors.New("elasticConfig not found")
		}

		{
			es, err := elasticsearch.NewClient(*client.Config.ElasticConfig)
			if err != nil {
				return err
			}
			client.Client = es
		}

		return nil
	}()

	return client, err
}

type Client struct {
	Config *Config
	Client *elasticsearch.Client
}

func (rec *Client) Do(request esapi.Request) (*Response, error) {
	var response *Response

	err := func() error {
		res, err := request.Do(context.Background(), rec.Client.Transport)
		if err != nil {
			return err
		}
		defer func() {
			_ = res.Body.Close()
		}()

		if res.IsError() {
			by := &bytes.Buffer{}
			_, err := io.Copy(by, res.Body)
			if err != nil {
				return errors.New("response error")
			}
			return errors.New(by.String())
		}

		response, err = NewResponse(res)
		if err != nil {
			return err
		}

		return nil
	}()

	return response, err
}

func (rec *Client) DoBulk(request *esapi.BulkRequest) (*Response, error) {
	var response *Response

	err := func() error {
		if request == nil {
			return errors.New("request is nil")
		}

		var err error

		response, err = rec.Do(request)
		if err != nil {
			return err
		}

		{
			var responseMap map[string]interface{}
			err = json.Unmarshal(response.Body, &responseMap)
			if err != nil {
				return err
			}

			if val, ok := responseMap["errors"]; ok {
				if val.(bool) {
					return errors.New(response.GetBodyString())
				}
			}
		}

		return nil
	}()

	return response, err
}

func (rec *Client) NewQueryBulk() (*QueryBulk, error) {
	return NewQueryBulk(rec)
}

package gostic

import (
	"github.com/elastic/go-elasticsearch/v7"
)

func NewClient(conf *Config) (*Client, error) {
	client := &Client{}
	err := client.Init(conf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewClientDefault() (*Client, error) {
	return NewClient(nil)
}

func NewConfig(elasticsearchConfig *elasticsearch.Config) *Config {
	config := &Config{}
	config.Init(elasticsearchConfig)
	return config
}

func NewConfigDefault() *Config {
	return NewConfig(nil)
}

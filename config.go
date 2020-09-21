package gostic

import "github.com/elastic/go-elasticsearch/v7"

func NewConfigDefault() *Config {
	return NewConfig(nil)
}

func NewConfig(elasticsearchConfig *elasticsearch.Config) *Config {
	var config = &Config{
		BulkLimit: 500,
	}

	if elasticsearchConfig == nil {
		elasticsearchConfig = &elasticsearch.Config{}
	}
	config.ElasticConfig = elasticsearchConfig

	return config
}

type Config struct {
	ElasticConfig *elasticsearch.Config
	BulkLimit     int
}

func (rec *Config) SetAddress(addressList ...string) {
	rec.ElasticConfig.Addresses = addressList
}

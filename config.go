package gostic

import "github.com/elastic/go-elasticsearch/v7"

type Config struct {
	ElasticConfig *elasticsearch.Config
}

func (rec *Config) Init(elasticsearchConfig *elasticsearch.Config) {
	if elasticsearchConfig == nil {
		elasticsearchConfig = &elasticsearch.Config{}
	}
	rec.ElasticConfig = elasticsearchConfig
}

func (rec *Config) AddAddress(strList ...string) {
	for _, val := range strList {
		rec.ElasticConfig.Addresses = append(rec.ElasticConfig.Addresses, val)
	}
}

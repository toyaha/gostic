package gostic

import (
	"reflect"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
)

// func TestNewConfigDefault(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want *Config
// 	}{
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewConfigDefault(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewConfigDefault() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestNewConfig(t *testing.T) {
// 	type args struct {
// 		elasticsearchConfig *elasticsearch.Config
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *Config
// 	}{
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewConfig(tt.args.elasticsearchConfig); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestConfig_SetAddress(t *testing.T) {
	type fields struct {
		ElasticConfig *elasticsearch.Config
	}
	type args struct {
		addressList []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "success",
			fields: fields{
				ElasticConfig: &elasticsearch.Config{},
			},
			args: args{
				addressList: []string{"a", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &Config{
				ElasticConfig: tt.fields.ElasticConfig,
			}
			rec.SetAddress(tt.args.addressList...)
			if !reflect.DeepEqual(tt.args.addressList, tt.want) {
				t.Errorf("Config.SetAddress()\ngot  = %v\nwant = %v", tt.args.addressList, tt.want)
			}
		})
	}
}

package gostic

import (
	"reflect"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
)

// func TestNewConfigDefault(t *testing.T) {
// }

// func TestNewConfig(t *testing.T) {
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
			name: "ok",
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

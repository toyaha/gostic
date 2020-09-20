package gostic

import (
	"reflect"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// func TestNewClient(t *testing.T) {
// 	type args struct {
// 		addressList []string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *Client
// 		wantErr bool
// 	}{
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := NewClient(tt.args.addressList...)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewClient() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestClient_Do(t *testing.T) {
	es, err := testGetClient()
	if err != nil {
		t.Errorf("Client.Do() error = %v", err)
		return
	}

	type fields struct {
		Config *Config
		Client *elasticsearch.Client
	}
	type args struct {
		request esapi.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "ok get",
			fields: fields{
				Config: es.Config,
				Client: es.Client,
			},
			args: args{
				request: esapi.CatAliasesRequest{
					Name: []string{"sample"},
				},
			},
			want: &Response{
				StatusCode: 200,
				Body:       nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &Client{
				Config: tt.fields.Config,
				Client: tt.fields.Client,
			}
			got, err := rec.Do(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.StatusCode != tt.want.StatusCode {
				t.Errorf("Client.Do()\ncode = %v\nwant = %v", got.StatusCode, tt.want.StatusCode)
			}
			if !reflect.DeepEqual(got.Body, tt.want.Body) {
				t.Errorf("Client.Do()\nbody = %v\nwant = %v", string(got.Body), string(tt.want.Body))
			}
		})
	}
}

func TestClient_DoBulk(t *testing.T) {
	es, err := testGetClient()
	if err != nil {
		t.Errorf("Client.DoBulk() error = %v", err)
		return
	}

	type fields struct {
		Config *Config
		Client *elasticsearch.Client
	}
	type args struct {
		request *esapi.BulkRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Config: es.Config,
				Client: es.Client,
			},
			args: args{
				request: &esapi.BulkRequest{
					Index: "sample",
					Body: strings.NewReader(`{"create":{"_id":1}}
{"long":1}
`),
				},
			},
			want: &Response{
				StatusCode: 200,
			},
			wantErr: false,
		},
		{
			name: "ng",
			fields: fields{
				Config: es.Config,
				Client: es.Client,
			},
			args: args{
				request: &esapi.BulkRequest{
					Index: "sample",
					Body: strings.NewReader(`{"create":{"_id":1}}
{"long":1}
`),
				},
			},
			want: &Response{
				StatusCode: 200,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &Client{
				Config: tt.fields.Config,
				Client: tt.fields.Client,
			}
			got, err := rec.DoBulk(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DoBulk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.StatusCode != tt.want.StatusCode {
				t.Errorf("Client.DoBulk()\ncode = %v\nwant = %v", got.StatusCode, tt.want.StatusCode)
			}
		})
	}
}

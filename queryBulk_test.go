package gostic

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// func TestNewQueryBulk(t *testing.T) {
// }

func TestQueryBulk_do(t *testing.T) {
	query, err := testGetQueryBulk()
	if err != nil {
		t.Errorf("QueryBulk.do() error = %v", err)
		return
	}
	query.Request.Index = testIndexName

	type fields struct {
		Client    *Client
		Request   *esapi.BulkRequest
		ValueList []string
		BulkLimit int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Response
		want1   int
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Client:    query.Client,
				Request:   query.Request,
				ValueList: []string{`{"create":{"_id":2}`, `{"long":2}`},
				BulkLimit: query.BulkLimit,
			},
			want: &Response{
				StatusCode: 200,
			},
			want1:   2,
			wantErr: false,
		},
		{
			name: "ng",
			fields: fields{
				Client:    query.Client,
				Request:   query.Request,
				ValueList: []string{`{"create":{"_id":2}`, `{"long":2}`},
				BulkLimit: query.BulkLimit,
			},
			want: &Response{
				StatusCode: 200,
			},
			want1:   2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &QueryBulk{
				Client:    tt.fields.Client,
				Request:   tt.fields.Request,
				ValueList: tt.fields.ValueList,
				BulkLimit: tt.fields.BulkLimit,
			}
			got, got1, err := rec.do()
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryBulk.do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.StatusCode != tt.want.StatusCode {
				t.Errorf("QueryBulk.do()\ngot = %v\nwant = %v", got.StatusCode, tt.want.StatusCode)
			}
			if got1 != tt.want1 {
				t.Errorf("QueryBulk.do()\ngot1 = %v\nwant = %v", got1, tt.want1)
			}
		})
	}
}

// func TestQueryBulk_Do(t *testing.T) {
// }

// func TestQueryBulk_DoFinish(t *testing.T) {
// }

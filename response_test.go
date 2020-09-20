package gostic

import (
	"net/http"
	"reflect"
	"testing"
)

// func TestNewResponse(t *testing.T) {
// 	type args struct {
// 		response *esapi.Response
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *Response
// 		wantErr bool
// 	}{
// 		// Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := NewResponse(tt.args.response)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NewResponse() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewResponse() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestResponse_GetBody(t *testing.T) {
	type fields struct {
		Header     http.Header
		StatusCode int
		Body       []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "success",
			fields: fields{
				Body: []byte(`{"sample"}`),
			},
			want: []byte(`{"sample"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &Response{
				Header:     tt.fields.Header,
				StatusCode: tt.fields.StatusCode,
				Body:       tt.fields.Body,
			}
			if got := rec.GetBody(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.GetBody()\ngot  = %v\nwant = %v", got, tt.want)
			}
		})
	}
}

func TestResponse_GetBodyString(t *testing.T) {
	type fields struct {
		Header     http.Header
		StatusCode int
		Body       []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				Body: []byte(`{"sample"}`),
			},
			want: `{"sample"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &Response{
				Header:     tt.fields.Header,
				StatusCode: tt.fields.StatusCode,
				Body:       tt.fields.Body,
			}
			if got := rec.GetBodyString(); got != tt.want {
				t.Errorf("Response.GetBodyString()\ngot  = %v\nwant = %v", got, tt.want)
			}
		})
	}
}

func TestResponse_GetBodyMap(t *testing.T) {
	type fields struct {
		Header     http.Header
		StatusCode int
		Body       []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Body: []byte(`{"sample":"value"}`),
			},
			want: map[string]interface{}{
				"sample": "value",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &Response{
				Header:     tt.fields.Header,
				StatusCode: tt.fields.StatusCode,
				Body:       tt.fields.Body,
			}
			got, err := rec.GetBodyMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("Response.GetBodyMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.GetBodyMap()\ngot  = %v\nwant = %v", got, tt.want)
			}
		})
	}
}

func TestResponse_GetBodyStruct(t *testing.T) {
	type fields struct {
		Header     http.Header
		StatusCode int
		Body       []byte
	}
	type args struct {
		structPtr interface{}
	}
	type sampleStruct struct {
		Sample string `json:"sample"`
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		obj     sampleStruct
		want    sampleStruct
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Body: []byte(`{"sample":"value"}`),
			},
			args: args{},
			obj:  sampleStruct{},
			want: sampleStruct{
				Sample: "value",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &Response{
				Header:     tt.fields.Header,
				StatusCode: tt.fields.StatusCode,
				Body:       tt.fields.Body,
			}
			tt.args.structPtr = &tt.obj
			if err := rec.GetBodyStruct(tt.args.structPtr); (err != nil) != tt.wantErr {
				t.Errorf("Response.GetBodyStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.obj, tt.want) {
				t.Errorf("Response.GetBodyMap()\ngot  = %+v\nwant = %+v", tt.obj, tt.want)
			}
		})
	}
}

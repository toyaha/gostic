package gostic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func NewResponse(response *esapi.Response) (*Response, error) {
	var data = &Response{}

	err := func() error {
		if response == nil {
			return errors.New("response not found")
		}

		data.Header = response.Header
		data.StatusCode = response.StatusCode

		{
			buf := &bytes.Buffer{}
			_, err := io.Copy(buf, response.Body)
			if err != nil {
				return err
			}

			data.Body = buf.Bytes()
		}

		return nil
	}()

	return data, err
}

type Response struct {
	Header     http.Header
	StatusCode int
	Body       []byte
}

func (rec *Response) GetBody() []byte {
	return rec.Body
}

func (rec *Response) GetBodyString() string {
	return string(rec.Body)
}

func (rec *Response) GetBodyMap() (map[string]interface{}, error) {
	var responseMap map[string]interface{}
	err := json.Unmarshal(rec.Body, &responseMap)
	if err != nil {
		return nil, err
	}
	return responseMap, nil
}

func (rec *Response) GetBodyStruct(structPtr interface{}) error {
	// reader := bytes.NewReader(rec.Body)
	// err := json.NewDecoder(reader).Decode(&structPtr)
	// if err != nil {
	// 	return err
	// }
	return json.Unmarshal(rec.Body, structPtr)
}

func (rec *Response) IsError() bool {
	return rec.StatusCode > 299
}

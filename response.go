package gostic

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
	"net/http"
)

type Response struct {
	Header     *http.Header
	StatusCode *int
	Body       []byte
}

func (rec *Response) Init(response *esapi.Response) error {
	if response == nil {
		return errors.New("response not exist")
	}

	rec.Header = &response.Header
	rec.StatusCode = &response.StatusCode

	{
		buf := &bytes.Buffer{}
		_, err := io.Copy(buf, response.Body)
		if err != nil {
			return err
		}

		rec.Body = buf.Bytes()
	}

	return nil
}

func (rec *Response) GetBody() []byte {
	return rec.Body
}

func (rec *Response) GetBodyMap() (map[string]interface{}, error) {
	var responseMap map[string]interface{}
	err := json.Unmarshal(rec.Body, &responseMap)
	if err != nil {
		return nil, err
	}

	return responseMap, nil
}

func (rec *Response) GetBodyString() string {
	return string(rec.Body)
}

func (rec *Response) GetBodyStruct(structPtr interface{}) error {
	reader := bytes.NewReader(rec.Body)
	err := json.NewDecoder(reader).Decode(&structPtr)
	if err != nil {
		return err
	}
	return nil
}

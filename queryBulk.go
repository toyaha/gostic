package gostic

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func NewQueryBulk(client *Client) (*QueryBulk, error) {
	var query = &QueryBulk{}

	err := func() error {
		if client == nil {
			return errors.New("client is nil")
		}
		query.Client = client
		query.BulkLimit = client.Config.BulkLimit
		query.Request = &esapi.BulkRequest{}

		return nil
	}()

	return query, err
}

type QueryBulk struct {
	Client    *Client
	Request   *esapi.BulkRequest
	ValueList []string
	BulkLimit int
}

func (rec *QueryBulk) do() (*Response, int, error) {
	var res *Response
	var bulkCount int
	var err error

	err = func() error {
		if rec.Client == nil {
			return errors.New("client is nil")
		}
		if rec.Request == nil {
			return errors.New("request is nil")
		}

		bulkCount = len(rec.ValueList)

		{
			rec.Request.Body = strings.NewReader(strings.Join(rec.ValueList, "\n") + "\n")
			var tmpList []string
			rec.ValueList = tmpList
		}

		res, err = rec.Client.Do(rec.Request)
		if err != nil {
			return err
		}

		if res.IsError() {
			return errors.New(res.GetBodyString())
		}

		{
			var responseMap map[string]interface{}
			err = json.Unmarshal(res.Body, &responseMap)
			if err != nil {
				return err
			}

			if val, ok := responseMap["errors"]; ok {
				if val.(bool) {
					return errors.New(res.GetBodyString())
				}
			}
		}

		return nil
	}()

	return res, bulkCount, err
}

func (rec *QueryBulk) Do() (*Response, int, error) {
	var res *Response
	var bulkCount int
	var err error

	err = func() error {
		if len(rec.ValueList) < rec.BulkLimit {
			return nil
		}

		res, bulkCount, err = rec.do()
		if err != nil {
			return err
		}

		return nil
	}()

	return res, bulkCount, err
}

func (rec *QueryBulk) DoFinish() (*Response, int, error) {
	var res *Response
	var bulkCount int
	var err error

	err = func() error {
		if len(rec.ValueList) < 1 {
			return nil
		}

		res, bulkCount, err = rec.do()
		if err != nil {
			return err
		}

		return nil
	}()

	return res, bulkCount, err
}

func (rec *QueryBulk) AddValueString(valueList ...string) {
	rec.ValueList = append(rec.ValueList, valueList...)
}

package gostic

const (
	testEsUri     = "http://localhost:9200"
	testIndexName = "sample"
)

func testGetClient() (*Client, error) {
	var client *Client

	err := func() error {
		es, err := NewClient(nil)
		if err != nil {
			return err
		}
		es.Config.SetAddress(testEsUri)

		client = es

		return nil
	}()

	return client, err
}

func testGetQueryBulk() (*QueryBulk, error) {
	var query *QueryBulk

	err := func() error {
		es, err := NewClient(nil)
		if err != nil {
			return err
		}
		es.Config.SetAddress(testEsUri)

		query, err = NewQueryBulk(es)

		return nil
	}()

	return query, err
}

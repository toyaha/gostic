package gostic

const (
	testEsUri     = "http://localhost:9201"
	testIndexName = "sample"
)

func testGetClient() (*Client, error) {
	var client *Client

	err := func() error {
		conf := NewConfig(nil)

		conf.SetAddress(testEsUri)

		es, err := NewClient(conf)
		if err != nil {
			return err
		}

		client = es

		return nil
	}()

	return client, err
}

func testGetQueryBulk() (*QueryBulk, error) {
	var query *QueryBulk

	err := func() error {
		conf := NewConfig(nil)

		conf.SetAddress(testEsUri)

		es, err := NewClient(conf)
		if err != nil {
			return err
		}

		query, err = NewQueryBulk(es)
		if err != nil {
			return err
		}

		return nil
	}()

	return query, err
}

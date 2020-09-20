package gostic

const (
	testEsUri = "http://localhost:9200"
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

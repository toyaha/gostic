# gostic

``` go
conf := gostic.NewConfig(nil)

uri := "http://localhost:9200"
conf.AddAddress(uri)

es, err := gostic.NewClient(nil)
if err != nil {
    return nil, err
}

return es, nil
```
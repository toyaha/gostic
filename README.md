# gostic

``` go
es, err := gostic.NewClient(nil)
if err != nil {
    return nil, err
}
es.Config.AddAddress(config.EsUrl)

return es, nil
```
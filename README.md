# gostic

``` go
conf := gostic.NewConfigDefault()
conf.AddAddress(config.EsUrl)
es, err := gostic.NewClient(conf)
if err != nil {
    return nil, err
}
return es, nil
```
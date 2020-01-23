# gostic

``` go
conf := gostic.NewConfigDefault()
conf.AddAddress(config.EsUrl)
elastic, err := gostic.NewClient(conf)
if err != nil {
    return nil, err
}
return elastic, nil
```
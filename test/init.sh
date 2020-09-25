#! /bin/bash




URI=http://localhost:9201




curl -XDELETE ${URI}/sample?pretty


curl -XPUT -H "Content-Type: application/json" ${URI}/sample?pretty -d'
{
  "mappings": {
    "properties": {
      "keyword": {
        "type": "keyword",
        "null_value": "NULL"
      },
      "text": {
        "type": "text"
      },
      "long": {
        "type": "long"
      },
      "integer": {
        "type": "integer"
      },
      "short": {
        "type": "short"
      },
      "float": {
        "type": "float"
      },
      "date": {
        "type": "date",
        "format": "yyyy/MM/dd HH:mm:ss||yyyy/MM/dd||yyyy-MM-dd'\''T'\''HH:mm:ss.Sz||yyyy-MM-dd'\''T'\''HH:mm:ss.SSz||yyyy-MM-dd'\''T'\''HH:mm:ss.SSSz||yyyy-MM-dd'\''T'\''HH:mm:ss.SSSSz||yyyy-MM-dd'\''T'\''HH:mm:ss.SSSSSz||yyyy-MM-dd'\''T'\''HH:mm:ss.SSSSSSz||yyyy-MM-dd'\''T'\''HH:mm:ss.SSSSSSSz||yyyy-MM-dd'\''T'\''HH:mm:ss.SSSSSSSSz||yyyy-MM-dd'\''T'\''HH:mm:ss.SSSSSSSSSz||yyyy-MM-dd'\''T'\''HH:mm:ssz||yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||yyyyMMdd||epoch_millis"
      },
      "boolean": {
        "type": "boolean"
      },
      "nested": {
        "type": "nested",
        "include_in_parent": true,
        "properties": {
          "long": {
            "type": "long"
          },
          "keyword": {
            "type": "keyword"
          }
        }
      }
    }
  }
}'


curl -XGET ${URI}/_alias?pretty

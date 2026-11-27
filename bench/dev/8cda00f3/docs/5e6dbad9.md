+++
title = "Data source HTTP API "
description = "Grafana Data source HTTP API"
keywords = ["grafana", "http", "documentation", "api", "data source"]
aliases = ["/http_api/datasource/"]
type = "docs"
[menu.docs]
name = "Data source"
parent = "http_api"
+++


# Data source API

## Get all datasources

`GET /api/datasources`

**Example Request**:

```http
GET /api/datasources HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJfMiutpDmyzWV2J8S5UsZjMHQ5HtHGOVMwh8a3S64OEWJ7Pw
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

[
  {
    "id":1,
    "orgId":1,
    "name":"datasource_elastic",
    "type":"elasticsearch",
    "access":"proxy",
    "url":"http://mydatasource.com",
    "password":"",
    "user":"",
    "database":"grafana-dash",
    "basicAuth":false,
    "basicAuthUser":"",
    "basicAuthPassword":"",
    "isDefault":false,
    "jsonData":null
  }
]
```

## Get a single data sources by Id

`GET /api/datasources/:datasourceId`

**Example Request**:

```http
GET /api/datasources/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJpDdgisIT8UBo9eLOaPI3uIcHCiQcIMBe7nye6GjvLG4pdVa
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{
  "id":1,
  "orgId":1,
  "name":"test_datasource",
  "type":"graphite",
  "access":"proxy",
  "url":"http://mydatasource.com",
  "password":"",
  "user":"",
  "database":"",
  "basicAuth":false,
  "basicAuthUser":"",
  "basicAuthPassword":"",
  "isDefault":false,
  "jsonData":null
}
```

## Get a single data source by Name

`GET /api/datasources/name/:name`

**Example Request**:

```http
GET /api/datasources/name/test_datasource HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJlStressFPappjS3G0uASKIpurlhyrRyrxVW4MwHx2BDZHNG
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{
  "id":1,
  "orgId":1,
  "name":"test_datasource",
  "type":"graphite",
  "access":"proxy",
  "url":"http://mydatasource.com",
  "password":"",
  "user":"",
  "database":"",
  "basicAuth":false,
  "basicAuthUser":"",
  "basicAuthPassword":"",
  "isDefault":false,
  "jsonData":null
}
```

## Get data source Id by Name

`GET /api/datasources/id/:name`

**Example Request**:

```http
GET /api/datasources/id/test_datasource HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJeGzhmUd3XuA8CqUCpa9sLoUbcjjNOXSOElt2BkDE6ANOMsQ
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{
  "id":1
}
```

## Create data source

`POST /api/datasources`

**Example Graphite Request**:

```http
POST /api/datasources HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJzBdmwZ1alE4ZYUaE3u5Y7pZHNZUMRLUApFHiyT4rIX9kXbJ

{
  "name":"test_datasource",
  "type":"graphite",
  "url":"http://mydatasource.com",
  "access":"proxy",
  "basicAuth":false
}
```

**Example CloudWatch Request**:

```http
POST /api/datasources HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJlMuhzQmQ4xRvVGnaou252vTjyFWcEKWCzIXYPEFblNCUluD

{
  "name": "test_datasource",
  "type": "cloudwatch",
  "url": "http://monitoring.us-west-1.amazonaws.com",
  "access": "proxy",
  "jsonData": {
    "authType": "keys",
    "defaultRegion": "us-west-1"
  },
  "secureJsonData": {
    "accessKey": "Au7kROvePCM3XaklFf9u", //should not be encoded
    "secretKey": "gKOutBVcBLitPihxFZBbPXXcmsb6FTB8FOKn" //should be Base-64 encoded
  }
}
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{"id":1,"message":"Datasource added", "name": "test_datasource"}
```

## Update an existing data source

`PUT /api/datasources/:datasourceId`

**Example Request**:

```http
PUT /api/datasources/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJlJaredMuvH8ZrC4CoMi54aCsLeMyTD3OAK4bSK2IJO8vvkY

{
  "id":1,
  "orgId":1,
  "name":"test_datasource",
  "type":"graphite",
  "access":"proxy",
  "url":"http://mydatasource.com",
  "password":"",
  "user":"",
  "database":"",
  "basicAuth":true,
  "basicAuthUser":"basicuser",
  "basicAuthPassword":"aWd5eXV3emdp",
  "isDefault":false,
  "jsonData":null
}
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{"message":"Datasource updated", "id": 1, "name": "test_datasource"}
```

## Delete an existing data source by id

`DELETE /api/datasources/:datasourceId`

**Example Request**:

```http
DELETE /api/datasources/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJyMeoeygwwEGDLTZy3R8GLyNEtewHBOMsFS9rZCOOYWGUBZJ
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{"message":"Data source deleted"}
```

## Delete an existing data source by name

`DELETE /api/datasources/name/:datasourceName`

**Example Request**:

```http
DELETE /api/datasources/name/test_datasource HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJeVmvzFrFHD5oXIZazu1RzArVMjYzvODitGO3PE2NoZ6EHbN
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{"message":"Data source deleted"}
```

## Data source proxy calls

`GET /api/datasources/proxy/:datasourceId/*`

Proxies all calls to the actual datasource.

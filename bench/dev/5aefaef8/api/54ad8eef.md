  * Suggested client recovery behavior:
    * Increase the value of the timeout param and retry with exponential
backoff.

## Response Status Kind

Kubernetes will always return the `Status` kind from any API endpoint when an
error occurs. Clients SHOULD handle these types of objects when appropriate.

A `Status` kind will be returned by the API in two cases:
  * When an operation is not successful (i.e. when the server would return a non
2xx HTTP status code).
  * When a HTTP `DELETE` call is successful.

The status object is encoded as JSON and provided as the body of the response.
The status object contains fields for humans and machine consumers of the API to
get more detailed information for the cause of the failure. The information in
the status object supplements, but does not override, the HTTP status code's
meaning. When fields in the status object have the same meaning as generally
defined HTTP headers and that header is returned with the response, the header
should be considered as having higher priority.

**Example:**

```console
$ curl -v -k -H "Authorization: Bearer VkXHzf1JNoTlcnanD6oz9X9ivlfEsxTs" https://10.240.122.184:443/api/v1/namespaces/default/pods/grafana

> GET /api/v1/namespaces/default/pods/grafana HTTP/1.1
> User-Agent: curl/7.26.0
> Host: 10.240.122.184
> Accept: */*
> Authorization: Bearer UhBYot2EIsQmaflyE2sh3I3nsujVyeGc
>

< HTTP/1.1 404 Not Found
< Content-Type: application/json
< Date: Wed, 20 May 2015 18:10:42 GMT
< Content-Length: 232
<
{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {},
  "status": "Failure",
  "message": "pods \"grafana\" not found",
  "reason": "NotFound",
  "details": {
    "name": "grafana",
    "kind": "pods"
  },
  "code": 404
}
```

`status` field contains one of two possible values:
* `Success`
* `Failure`
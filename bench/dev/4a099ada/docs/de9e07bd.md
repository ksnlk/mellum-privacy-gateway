---
title: 'Auth Session'
apiurl: '/api/v1/user/auth_session'
type: 'GET'
layout: default
---

透过Session检查去判定使用者可否存取资源

### RequestHeader
透过RequestHeader 的 Apitoken做验证
```"RequestHeader": {
  "Apitoken": "{\"bqum\":\"bvde\",\"xax\":\"461x4952k71064h37eea8315gc631727\"}",
  "X-Forwarded-For": " 127.0.0.1"
}```

### Response

Session 为有效
```Status: 200```
```{"message":"session is vaild!"}```

For errors responses, see the [response status codes documentation](#/response-status-codes).

# echo-headers

A simple tool for verifying request headers. As the name implies, echo-headers will repond with the set of request headers it received, including anything sensitive like Authorization, etc. Useful for troubleshooting situations where headers are being mutated/changed and you want to verify exactly what's being received.

#### Start the server
```
➜ go run main.go
Server started on :8080
```

#### Send a request to the server
```
➜ curl http://localhost:8080 -H "Authorization: big secret" -H 'content-type: text/plain;charset=utf-8'
Content-Type: text/plain;charset=utf-8
User-Agent: curl/7.88.1
Accept: */*
Authorization: big secret
```

Note that we're not passing `-i` or `-I` to curl in this example. The response is the request's headers written to the response body.

## Security Considerations

All headers and their values are displayed in the response to the client, including any security or auth headers that may be present. Importantly, the server _doesn't_ log headers on its end so they won't inadvertently be persisted to disk or your log aggregation system.
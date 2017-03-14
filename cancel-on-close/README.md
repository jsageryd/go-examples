# Cancel on close

Middleware to cancel the request context if the client disconnects before the
request completes.

## Usage
Start the server:
```
go run cancel-on-close.go
```

Make a request and cancel (ctrl-c) it within 5 seconds:
```
curl :8080/
(ctrl-c)
```

# gomodsmpl
go 1.16 の module mode で作ってみる

```shell
$ go run main.go post nashiki 30 
```

```shell
$ go build users/users.go

$ curl -X POST 127.0.0.1:1323/user -H "Content-Type: application/json" -d '{"name":"nashiki","age": 30}' 
"nashiki"
```

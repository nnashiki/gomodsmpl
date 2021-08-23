# gomodsmpl
go 1.16 の module mode で作ってみる

# ファイル単体の実行

```shell
$ go run main.go post nashiki 30 
```

# api server

```shell
$ go build users/users.go

$ curl -X POST 127.0.0.1:1323/user -H "Content-Type: application/json" -d '{"name":"nashiki","age": 30}' 
"nashiki"
```

# orm

```shell
$ go run gormmysql.go
```

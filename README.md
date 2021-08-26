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

## gorm

- doc
  - https://github.com/go-gorm/mysql
  - https://gorm.io/ja_JP/docs/index.html
- メソッドが参考になる
  - https://qiita.com/stranger1989/items/914bef8cc6ad77ef8350
  - https://selegee.com/7230/

## SQLBoiler

- https://zenn.dev/gami/articles/0fb2cf8b36aa09#orm%E3%81%A8%E3%81%AF%E3%81%A9%E3%81%86%E3%81%84%E3%81%86%E6%8A%80%E8%A1%93%E3%81%8B

## Go の良いところ
- https://qiita.com/methane/items/5ad7c092c0d426db4ab5
- https://qiita.com/methane/items/b627f20457873a504638
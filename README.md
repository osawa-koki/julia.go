# julia.go

🦊🦊🦊 Go言語でジュリア集合を描写する。  
普通に描写と、goroutineを使用した2つを用意。  

Goは偉大だ🐧🐧🐧  

![成果物](./.development/img/fruit.png)  

## 実行方法

DevContainer内に入り、以下のコマンドを実行。  

```shell
# 通常版
go run ./normal/main.go

# goroutineを使った最速版
go run ./faster/main.go
```

しっかりとビルド、実行するには、、、  

```shell
# 通常版
go build -o ./bin/normal ./normal/main.go
./bin/normal

# goroutineを使った最速版
go build -o ./bin/faster ./normal/main.go
./bin/faster
```

# clichat_client

## このリポジトリについて

clichatは、Goで制作したCLI上でチャットが出来るアプリケーションです。

ClientとServerに分かれており、それぞれが独立したコンテナとして動作をしながら同一のDockerネットワークで通信をします。

![overview](/Users/Ryusei/go/src/github.com/RyuseiNomi/clichat_client/assets/overview.png)

このリポジトリでは、クライアントコンテナを管理するdocker-composeファイル、そしてメッセージの送受信を行うソースコードを管理しています。



## ファイル構造

```
.
├── Dockerfile
├── Makefile
├── README.md
├── clichat_client
├── docker-compose.yml
├── main.go
└── tracer
    ├── tracer.go
    └── tracer_test.go
```

[補足] tracerパッケージは、[Go言語によるWebアプリケーション開発](https://www.oreilly.co.jp/books/9784873117522/)を参考に開発



## アプリの立ち上げ

1. [こちらのリポジトリ](https://github.com/RyuseiNomi/clichat_goserver)よりServerコンテナを立ち上げる
2. `$docker-compose up-d`を実行し、２つのコンテナが立ち上がることを確認する
3. ターミナルで画面分割を行うなどして、それぞれのクライアントコンテナに入る
4. チャットをお楽しみください
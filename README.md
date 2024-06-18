# hackathon_backend
ハッカソンのバックエンド用リポジトリ

## バックエンドの構造
主な構成要素は以下の通りです。

- | -api | openAPIの定義ファイルが格納されています。 |
- | -cmd | バックエンドを起動する時のコマンドが格納されています。 |
- | -init_query | バックエンドの初期化時に実行するクエリが格納されています。 |
- | -internal | バックエンドの内部処理が格納されています。 |
- | -test | バックエンドのテストが格納されています。 |

## バックエンドの起動方法
以下の手順でバックエンドをローカルで起動できます。 
```bash
make dev-up
```

## バックエンドのテスト方法
以下の手順でバックエンドのテストを実行できます。
```bash
make test
```

## バックエンドの停止方法
以下の手順でバックエンドを停止できます。
```bash
make dev-down
```

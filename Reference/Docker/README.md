# Docker備忘

## ビルド
```zsh
# 中間レイヤーのキャッシュを使用してビルドするため、高速
docker-compose build
```

```zsh
# キャッシュを無効にするため、確実に最新の変更や依存関係を含んだイメージを作成できる
docker-compose build --no-cache
```

## コンテナ起動
```zsh
# サービスをフォアグラウンド(ターミナル上)で起動
# 開発時やデバッグ時にサービスの出力を確認できる
docker-compose up
```

```zsh
# サービスをバックグラウンドで起動
# docker-compose logsで後からログを確認できる
docker-compose up -d
```

## 実行
### コンテナ内のターミナルに入って実行

```zsh
docker-compose exec -it <container-name or container-id> /bin/sh
```

## 後片付け
### 停止と削除（コンテナ・ネットワーク）
```zsh
docker-compose down
```


### 停止と削除（コンテナ・ネットワーク・イメージ）
```zsh
docker-compose down --rmi all
```

### 停止と削除（コンテナ・ネットワーク・ボリューム）
```zsh
docker-compose down -v
```

### 全てを無に返す
```zsh
docker-compose down --rmi all --volumes
```
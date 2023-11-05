# go-sample
- Goの文法を理解するためのリポジトリ
- 備忘系は[Reference](./Reference/)に記載

# 作成背景
- 会社の勉強会でGoを使用したWebアプリ開発を行っており、基本的な文法を学ぼうと思い作成。
- また、普段あまりDockerを使用しないので、Dockerの利用方法の備忘も兼ねて作成。

# ブランチ運用
- ブランチを切る際はissueと紐づけるため、suffixに"#issue番号"を記載する
```zsh
git checkout -b <ブランチ名>_#issue番号
```

- コミットの際はissueと紐づけるため、コミットメッセージに"#issue番号"を記載する
```zsh
git commit -m "コミットメッセージ #issue番号"
```

# How to use
1. クローン
```zsh
https://github.com/YuyaNakamura0139/go-sample.git
```

2. コンテナ作成
```zsh
docker-compose build
```

3. コンテナ起動
```zsh
docker-compose up
```

3. コンテナ内のターミナルに入る
```zsh
docker-compose exec -it go-sample /bin/sh
```

4. コンテナ停止
```zsh
docker-compose stop
```

5. コンテナ再起動
```zsh
docker-compose start
```

6. コンテナ削除
```zsh
docker-compose down
```

7. 全部削除したい場合
```zsh
docker-compose down --rmi all --volumes
```

# 参考文献
- https://zenn.dev/yusuke49/articles/9ed37838861b1d
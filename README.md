# 起動

```
$ docker-compose build
$ docker-compose up -d # バックエンドで実行
$ docker-compose up # そのまま実行
```

# ログの確認（バックエンドでの実行時）

```
$ docker-compose logs golang # go
$ docker-compose logs mysql # mysql
```

# ディレクトリ構成

- app: アプリケーションフォルダ
  - controller:
  - database: 全体的なデータベース周りの処理
  - middleware:
  - models: モデル定義など
  - main.go: アプリケーション起動とルーティング
- docker: docker 関連
  - api: api 関連
  - db: データベース関連
    - db_data: データベースのデータを保持
    - db_init: データベース初期化用の SQL

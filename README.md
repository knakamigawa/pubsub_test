# pubsub検証アプリケーション

## GCPプロジェクトの作成

## GCPでトピック作成

## publisherの構築

コマンドライン版とWebインターフェース版
コマンドライン版は.envrcを設定後コマンドを実行することでtopicへpublishを行います
```
go run main_cmd.go
```

## subscriberの構築

```
// m1 mac
docker buildx build --platform linux/amd64 --load ./ --tag gcr.io/____youre_project____/____youre_application_name____
// コンテナイメージをGCRにpush
docker push gcr.io/____youre_project____/____youre_application_name____
```

上記でpushしたイメージでCloudRunのアプリケーションを作成

## 作成したトピックのサブスクライバを作成

作成したCloudRunをトリガするPushで作成
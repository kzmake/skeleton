# skeleton: Golang用テンプレートプロジェクト

## テンプレートを利用してプロジェクトを始める
### プロジェクトのパス変更

下記のコマンドを初回のみ実行してください。

```bash
./init.sh {{ user_name }}/{{ project_name }}

# ex:
#   ./init.sh hogehoge/fugafuga
```

### マイクロサービスの追加

1. `api/{{ マイクロサービス名 }}/v1` にて.protoの追加
1. `backend/svc/{{ マイクロサービス名 }}` にてgRPCアプリケーションの追加
1. `Makefile` にて SERVICES 変数に `backend/svc/{{ マイクロサービス名 }}` を設定
1. `manifests/base/svc-{{ マイクロサービス名 }}.yaml` としてマニフェストの追加
1. `skaffold.app.yaml` にて追加したマイクロサービスのコンテナイメージのビルドを追加

## Requirements

このプロジェクトでは、

- [buf](https://docs.buf.build/)

などを使用します。初回のみ

```bash
make install
```

で必要なツールをインストールしてください。

## for developer
### ProtocolBuffersからコードを生成＆フォーマットやLinterを実施する

```bash
make
```

### ローカル(kind)環境を準備する

```bash
make kind
```

### アプリケーションをビルド＆デプロイする

```bash
make dev
```

### ローカル環境(kind)を削除する

```bash
make clean
```

## Production環境へのデプロイ

下記のコマンドを実施することでデプロイできます。

```bash
make deploy-production
```

もし、デプロイ済みのアプリケーションを削除する場合は

```bash
make destroy-production
```

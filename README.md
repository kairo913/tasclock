[README in English](https://github.com/kairo913/tasclock/blob/main/README-en.md)

# TasClock: タスク管理と時給計算を融合したアプリ

TasClock は、タスク管理とタイムトラッキングを融合したアプリです。各タスクに時間を計測し、設定した単価から現在の時給を計算します。

## 主な機能

-   タスクの作成、編集、削除
-   各タスクの開始時間と終了時間の記録
-   設定した単価に基づいて、現在の時給をリアルタイムで計算
-   タスクごとの作業時間の履歴をグラフで表示
-   目標時給を設定して、モチベーションを維持
-   CSV 形式でデータのエクスポート

## 依存関係

Go 1.18+  
NPM (Node 15+)

## Wails のインストール

`go install github.com/wailsapp/wails/v2/cmd/wails@latest` を実行して、Wails CLI をインストールしてください。

## 開発モード

開発モードで実行するには、プロジェクトディレクトリで `wails dev` を実行してください。これにより Vite 開発用サーバーが実行され、フロントエンドの変更を非常に高速にホットリロードできるようになります。ブラウザで開発し、Go メソッドにアクセスしたい場合は、http://localhost:34115 にアクセスしてください。ブラウザでこのサーバーに接続すると、devtoolsから Go のコードを呼び出すことができます。

## ビルド

再配布可能な本番環境用パッケージをビルドするには、 `wails build` を使用してください。  
そうすることで、プロジェクトがコンパイルされ、`build/bin` ディレクトリ内に本番配布用のバイナリが出力されます。

## 参考

[Wails: インストール](https://wails.io/ja/docs/gettingstarted/installation)

[Wails: プロジェクトのコンパイル](https://wails.io/ja/docs/gettingstarted/building)
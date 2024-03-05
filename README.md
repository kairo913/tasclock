[README in English](https://github.com/kairo913/tasclock/README-en.md)

# TasClock: タスク管理と時給計算を融合したアプリ

TasClock は、タスク管理とタイムトラッキングを融合したアプリです。各タスクに時間を計測し、設定した単価から現在の時給を計算します。

## 主な機能

-   タスクの作成、編集、削除
-   各タスクの開始時間と終了時間の記録
-   設定した単価に基づいて、現在の時給をリアルタイムで計算
-   タスクごとの作業時間の履歴をグラフで表示
-   目標時給を設定して、モチベーションを維持
-   CSV 形式でデータのエクスポート

## ビルド方法

### 依存関係

Go 1.18+  
NPM (Node 15+)

### Wails のインストール

`go install github.com/wailsapp/wails/v2/cmd/wails@latest`

### ビルド

`wails build`
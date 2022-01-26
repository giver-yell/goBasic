# goBasic

## Setting

`go get golang.org/x/tools/cmd/godoc`

- `fmt`コマンド確認  
  `go doc fmt`

## Test

- テスト実行
  `go test -v ./...`

## third party の package インストール

- 株価分析の package
  `go get github.com/markcheno/go-talib`
- semaphore
  `go mod download golang.org/x/sync`
  `go install golang.org/x/sync/semaphore`
- go ini
  `go get gopkg.in/ini.v1`
- JSON-RPC 2.0 over WebSocket で Bitcoin の価格をリアルタイムに取得する
  `go get github.com/gorilla/websocket`

## ドキュメントインストール

- golang
  `go get golang.org/x/tools/cmd/godoc`

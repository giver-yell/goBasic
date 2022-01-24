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

## ドキュメントインストール

- golang
  `go get golang.org/x/tools/cmd/godoc`

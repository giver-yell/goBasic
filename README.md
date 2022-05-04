# goBasic

## sqlite3 操作

- ログイン  
  `sqlite3 example.sql`
- テーブル確認  
  sqlite> `.table`

## Test

- テスト実行  
  `go test -v ./...`

## Setting

### third party の package インストール

- 株価分析の package  
  `go get github.com/markcheno/go-talib`
- semaphore  
  `go mod download golang.org/x/sync`  
  `go install golang.org/x/sync/semaphore`
- go ini  
  `go get gopkg.in/ini.v1`
- JSON-RPC 2.0 over WebSocket で Bitcoin の価格をリアルタイムに取得する  
  `go get github.com/gorilla/websocket`

### sqlite3 インストール

- sqlite3 インストール  
  ` brew install sqlite`
- xcode インストール  
  - ホームページよりインストール https://developer.apple.com/download/all/?q=xcode
- xcode コマンドツールをインストール  
  `xcode-select --install`
- gcc のインストール確認  
  `gcc --version`
- sqlite3 のサードパーティのパッケージインストール  
  `go get github.com/mattn/go-sqlite3`

### ドキュメントインストール

- golang  
  `go get golang.org/x/tools/cmd/godoc`

- godocをコマンドで確認
   - 例)`fmt`を確認      
  `go doc fmt`  
  
### go1.18インストール
1. go1.18をofficialのリファレンスからダウンロード  
https://go.dev/dl/  
2. `.zshrc`にPATHを追加  
```
path=(
  /usr/local/go/bin
)
```

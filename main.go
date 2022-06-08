package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// 77.iniでConfigの設定ファイルを読み込む
type ConfigList struct {
	Port      int
	DBName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		// MustInt(): 空なら0が入る
		Port:   cfg.Section("web").Key("port").MustInt(),
		DBName: cfg.Section("db").Key("name").MustString("exampe.sql"),
		// String(): 空なら''が入る
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DBName, Config.DBName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}

// 76.Semaphore
// 同時にgoroutineが実行できる数を指定
// var s *semaphore.Weighted = semaphore.NewWeighted(1)

// func longProcess(ctx context.Context) {
// 	// 処理待ちのキャンセル
// 	isAcquire := s.TryAcquire(1)
// 	if !isAcquire {
// 		fmt.Println("Could not get lock")
// 		return
// 	}
// 	// ロック
// 	// if err := s.Acquire(ctx, 1); err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// リリース
// 	defer s.Release(1)

// 	fmt.Println("Wait...")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Done")
// }

// func main() {
// 	ctx := context.TODO()
// 	go longProcess(ctx)
// 	go longProcess(ctx)
// 	go longProcess(ctx)
// 	time.Sleep(2 * time.Second)
// 	go longProcess(ctx)
// 	time.Sleep(5 * time.Second)
// }

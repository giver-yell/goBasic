package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	"log"
)

// 78.DB操作
var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	// テーブル作成
	cmd := `CREATE TABLE IF NOT EXISTS person(
				name STRING,
				age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// insert
	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err = DbConnection.Exec(cmd, "Nancy", 20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// update
	// cmd = "UPDATE person SET age = ? WHERE name = ? "
	// _, err = DbConnection.Exec(cmd, 25, "Mike")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// multi select
	// cmd = "SELECT * FROM person"
	// rows, _ := DbConnection.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// // まとめてエラー取得
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// single select
	// cmd = "SELECT * FROM person WHERE age = ?"
	// row := DbConnection.QueryRow(cmd, 100)
	// var p Person
	// err = row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// delete
	// cmd = "DELETE FROM person WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// テーブル名を動的に変更
	// テーブル名はプレースホルダが使えない(本来であればSQLインジェクションの対策が必要)
	tableName := "person"
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	// まとめてエラー取得
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}

// 77.iniでConfigの設定ファイルを読み込む
// type ConfigList struct {
// 	Port      int
// 	DBName    string
// 	SQLDriver string
// }

// var Config ConfigList

// func init() {
// 	cfg, _ := ini.Load("config.ini")
// 	Config = ConfigList{
// 		// MustInt(): 空なら0が入る
// 		Port:   cfg.Section("web").Key("port").MustInt(),
// 		DBName: cfg.Section("db").Key("name").MustString("exampe.sql"),
// 		// String(): 空なら''が入る
// 		SQLDriver: cfg.Section("db").Key("driver").String(),
// 	}
// }

// func main() {
// 	fmt.Printf("%T %v\n", Config.Port, Config.Port)
// 	fmt.Printf("%T %v\n", Config.DBName, Config.DBName)
// 	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
// }

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

/*
セクション13
*/
package main

import (
	"fmt"
	"io/ioutil"
)

// 85.Web Applications 1 -ioutil
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// 0600は読み書き可能
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "test", Body: []byte("This is a simple page")}
	p1.save()

	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))
}

/*
セクション12
*/
// package main

// 84.DB操作
// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// var DbConnection *sql.DB

// type Person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	// DB接続
// 	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
// 	defer DbConnection.Close()

// 	// CREATE TABLE
// 	cmd := `CREATE TABLE IF NOT EXISTS person(
// 		name STRING,
// 		age INT)`
// 	_, err := DbConnection.Exec(cmd)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// INSERT
// 	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
// 	// execの結果は返さない
// 	// _, err = DbConnection.Exec(cmd, "Nancy", 20)
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }

// 	// UPDATE
// 	// cmd = "UPDATE person SET age = ? WHERE name = ?"
// 	// _, err = DbConnection.Exec(cmd, 25, "Mike")
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }

// 	// マルチSELECT
// 	// cmd = "SELECT * FROM person"
// 	// rows, _ := DbConnection.Query(cmd)
// 	// defer rows.Close()
// 	// var pp []Person
// 	// for rows.Next() {
// 	// 	var p Person
// 	// 	err := rows.Scan(&p.Name, &p.Age)
// 	// 	if err != nil {
// 	// 		log.Println(err)
// 	// 	}
// 	// 	pp = append(pp, p)
// 	// }
// 	// err = rows.Err()
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }
// 	// for _, p := range pp {
// 	// 	fmt.Println(p.Name, p.Age)
// 	// }

// 	// single select
// 	// cmd = "SELECT * FROM person WHERE age = ?"
// 	// row := DbConnection.QueryRow(cmd, 20)
// 	// var p Person
// 	// err = row.Scan(&p.Name, &p.Age)
// 	// if err != nil {
// 	// 	if err == sql.ErrNoRows {
// 	// 		log.Println("No rows")
// 	// 	} else {
// 	// 		log.Println(err)
// 	// 	}
// 	// }
// 	// fmt.Println(p.Name, p.Age)

// 	// DELETE
// 	// cmd = "DELETE FROM person WHERE age = ?"
// 	// _, err = DbConnection.Exec(cmd, 20)
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// }

// 	// table nameを動的にselect
// 	tableName := "person"
// 	cmd = fmt.Sprintf("SELECT * FROM  %s", tableName)
// 	rows, _ := DbConnection.Query(cmd)
// 	defer rows.Close()
// 	var pp []Person
// 	for rows.Next() {
// 		var p Person
// 		err := rows.Scan(&p.Name, &p.Age)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		pp = append(pp, p)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	for _, p := range pp {
// 		fmt.Println(p.Name, p.Age)
// 	}
// }

/*
セクション11
*/
// 80.JSON-RPC 2.0 over WebSocketでBitcoinの価格をリアルタイムに取得する
// import (
// 	"log"
// 	"net/url"

// 	"github.com/gorilla/websocket"
// )

// type JsonRPC2 struct {
// 	Version string      `json:"jsonrpc"`
// 	Method  string      `json:"method"`
// 	Params  interface{} `json:"params"`
// 	Result  interface{} `json:"result,omitempty"`
// 	Id      *int        `json:"id,omitempty"`
// }
// type SubscribeParams struct {
// 	Channel string `json:"channel"`
// }

// func main() {
// 	u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path: "/json-rpc"}
// 	log.Printf("connecting to %s", u.String())

// 	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		log.Fatal("dial:", err)
// 	}
// 	defer c.Close()

// 	if err := c.WriteJSON(&JsonRPC2{Version: "2.0", Method: "subscribe", Params: &SubscribeParams{"lightning_ticker_BTC_JPY"}}); err != nil {
// 		log.Fatal("subscribe:", err)
// 		return
// 	}

// 	for {
// 		message := new(JsonRPC2)
// 		if err := c.ReadJSON(message); err != nil {
// 			log.Println("read:", err)
// 			return
// 		}

// 		if message.Method == "channelMessage" {
// 			log.Println(message.Params)
// 		}
// 	}
// }

// 77.iniでconfigファイルの設定を読み込む
// import (
// 	"fmt"

// 	"gopkg.in/ini.v1"
// )
// type ConfigList struct {
// 	Port      int
// 	Dbname    string
// 	SQLDriver string
// }

// var Config ConfigList

// func init() {
// 	cfg, _ := ini.Load("config.ini")
// 	Config = ConfigList{
// 		Port:      cfg.Section("web").Key("port").MustInt(),
// 		Dbname:    cfg.Section("db").Key("name").MustString("example.sql"),
// 		SQLDriver: cfg.Section("db").Key("sqllite3").String(),
// 	}
// }

// func main() {
// 	fmt.Printf("%T %v\n", Config.Port, Config.Port)
// 	fmt.Printf("%T %v\n", Config.Dbname, Config.Dbname)
// 	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
// }

// 76.semaphore
// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"golang.org/x/sync/semaphore"
// )

// // 同時実行数を指定
// var s *semaphore.Weighted = semaphore.NewWeighted(2)

// func longProcess(ctx context.Context) {
// 	// 他の処理をキャンセルする
// 	isAcquire := s.TryAcquire(1)
// 	if !isAcquire {
// 		fmt.Println("Could not get lock")
// 		return
// 	}

// 	// blocking(他の処理を待たせる)
// 	// if err := s.Acquire(ctx, 1); err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	defer s.Release(1)
// 	fmt.Println("wait ...")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Done")
// }

// func main() {
// 	ctx := context.TODO()
// 	go longProcess(ctx)
// 	go longProcess(ctx)
// 	go longProcess(ctx)
// 	time.Sleep((5 * time.Second))
// }

/*
セクション10
*/
// 75.hmacでAPI認証
// import (
// 	"crypto/hmac"
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"fmt"
// )

// var DB = map[string]string{
// 	"User1Key": "User1Secret",
// 	"User2Key": "User2Secret",
// }

// func Server(apiKey, sign string, data []byte) {
// 	apiSecret := DB[apiKey]
// 	h := hmac.New(sha256.New, []byte(apiSecret))
// 	h.Write(data)
// 	expectedHMAC := hex.EncodeToString(h.Sum(nil))
// 	fmt.Println(sign == expectedHMAC)
// }

// func main() {
// 	const apiKey = "User2Key"
// 	const apiSecret = "User2Secret"

// 	data := []byte("data")
// 	h := hmac.New(sha256.New, []byte(apiSecret))
// 	h.Write(data)
// 	sign := hex.EncodeToString(h.Sum(nil))

// 	fmt.Println(sign)

// 	Server(apiKey, sign, data)
// }

// 74. json.UnmarshalとMarshalとエンコード
// import (
// 	"encoding/json"
// 	"fmt"
// )

// type T struct{}

// type Person struct {
// 	// jsonで隠す
// 	// Name string `json:"-"`
// 	Name string `json:"name,omitempty"`
// 	// jsonでstringへ変換
// 	// Age       int      `json:"age,string"`
// 	// 0の場合隠す
// 	Age       int      `json:"age,omitempty"`
// 	Nicknames []string `json:"nicknames"`
// 	// structを隠す場合はポインタ
// 	T *T `json:"T,omitempty"`
// }

// // unmarshal変換
// func (p *Person) UnmarshalJSON(b []byte) error {
// 	type Person2 struct {
// 		Name string
// 	}
// 	var p2 Person2
// 	err := json.Unmarshal(b, &p2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	p.Name = p2.Name + "!"
// 	return err
// }

// // JSONの値変換
// // func (p Person) MarshalJSON() ([]byte, error) {
// // 	v, err := json.Marshal(&struct {
// // 		Name string
// // 	}{
// // 		Name: "Mr." + p.Name,
// // 	})
// // 	return v, err
// // }

// func main() {
// 	// b := []byte(`{"name":"Mike","age":"20","nicknames":["a","b","c"]}`)
// 	// b := []byte(`{"name":"","age":0,"nicknames":["a","b","c"]}`)
// 	b := []byte(`{"name":"Mike","age":0,"nicknames":["a","b","c"]}`)

// 	var p Person
// 	if err := json.Unmarshal(b, &p); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(p.Name, p.Age, p.Nicknames)

// 	v, _ := json.Marshal(p)
// 	fmt.Println(string(v))
// }

// 73.http
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// )

// func main() {
// 	// resp, _ := http.Get("http://example.com")
// 	// defer resp.Body.Close()
// 	// body, _ := ioutil.ReadAll(resp.Body)
// 	// fmt.Println(string(body))

// 	base, _ := url.Parse("http://example.com")
// 	reference, _ := url.Parse("/test?a=1&b=2")
// 	endpoint := base.ResolveReference(reference).String()
// 	fmt.Println(endpoint)
// 	// get
// 	req, _ := http.NewRequest("GET", endpoint, nil)
// 	// post
// 	// req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
// 	// header追加
// 	req.Header.Add("If-None-Match", `W/"wyzzy"`)
// 	// urlクエリ追加
// 	q := req.URL.Query()
// 	q.Add("c", "3&%")
// 	fmt.Println(q)
// 	fmt.Println(q.Encode())

// 	var client *http.Client = &http.Client{}
// 	resp, _ := client.Do(req)
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(body))
// }

/*
セクション9
*/
// 72.ioutil
// import (
// 	"bytes"
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {
// content, err := ioutil.ReadFile("main.go")
// if err != nil {
// 	log.Fatalln(err)
// }
// fmt.Println(string(content))

//  ファイル出力
// if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
// 	log.Fatalln(err)
// }

// byte出力
// 	r := bytes.NewBuffer([]byte("abc"))
// 	content, _ := ioutil.ReadAll(r)
// 	fmt.Println(string(content))
// }

// 71.context
// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func longProcess(ctx context.Context, ch chan string) {
// 	fmt.Println("run")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("finish")
// 	ch <- "result"
// }

// func main() {
// 	ch := make(chan string)
// 	ctx := context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
// 	defer cancel()
// 	// 何も処理したくない場合
// 	// ctx := context.TODO()
// 	go longProcess(ctx, ch)
// 	// longProcessを実行させたくない場合
// 	// cancel()

// CTXLOOP:
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println(ctx.Err())
// 			break CTXLOOP
// 		case <-ch:
// 			fmt.Println("success")
// 			break CTXLOOP
// 		}
// 	}
// 	fmt.Println("#######")
// }

// 70.iota
// const (
// 	c1 = iota
// 	c2
// 	c3
// )

// const (
// 	_      = iota
// 	KB int = 1 << (10 * iota)
// 	MB
// 	GB
// )

// func main() {
// 	fmt.Println(c1, c2, c3)
// 	fmt.Println(KB, MB, GB)
// }

// 69.sort
// func main() {
// 	i := []int{1, 2, 3, 4, 5}
// 	s := []string{"d", "a", "f"}
// 	p := []struct {
// 		Name string
// 		Age  int
// 	}{
// 		{"Nancy", 20},
// 		{"Vera", 40},
// 		{"Mike", 30},
// 		{"Bob", 50},
// 	}
// 	fmt.Println(i, s, p)
// 	sort.Ints(i)
// 	sort.Strings(s)
// 	// for で回すイメージ
// 	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
// 	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })

// 	fmt.Println(i, s, p)
// }

// 68.regex
// func main() {
// 	match, _ := regexp.MatchString("a[a-z]+e", "apple")
// 	fmt.Println(match)

// 	r := regexp.MustCompile("a[a-z]+e")
// 	ms := r.MatchString("apple")
// 	fmt.Println(ms)

// 	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
// 	fs := r2.FindString("/view/test")
// 	fmt.Println(fs)
// 	fss := r2.FindStringSubmatch("/view/test")
// 	fmt.Println(fss, fss[0], fss[1], fss[2])
// 	fss = r2.FindStringSubmatch("/edit/test")
// 	fmt.Println(fss, fss[0], fss[1], fss[2])
// 	fss = r2.FindStringSubmatch("/save/test")
// 	fmt.Println(fss, fss[0], fss[1], fss[2])

// }

// 67.time
// postgresのtime の型は RFC3339
// func main() {
// 	t := time.Now()
// 	fmt.Println(t)
// 	fmt.Println(t.Format(time.RFC3339))
// 	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
// }

/*
セクション8
*/

// 64.サードパーティのpackagインストール
// import (
// 	"fmt"

// 	"github.com/markcheno/go-quote"
// 	"github.com/markcheno/go-talib"
// )

// func main() {
// 	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
// 	fmt.Print(spy.CSV())
// 	rsi2 := talib.Rsi(spy.Close, 2)
// 	fmt.Println(rsi2)
// }

// 60.パッケージ
// import (
// 	"fmt"

// 	"github.com/giver-yell/goBasic/mylib"
// 	"github.com/giver-yell/goBasic/mylib/under"
// )

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	fmt.Println(mylib.Average(s))

// 	mylib.Say()
// 	under.Hello()
// 	person := mylib.Person{Name: "Mike", Age: 20}
// 	fmt.Println(person)
// }

// func main() {
// 	totalScore := 0
// 	// 引数にtotalScoreのポインタを渡してください
// 	ask(1, "dog", &totalScore)
// 	ask(2, "cat", &totalScore)
// 	ask(3, "fish", &totalScore)

// 	fmt.Println("スコア", totalScore)
// }

// // 渡されるtotalScoreのポインタを受け取るように変更してください
// func ask(number int, question string, scorePtr *int) {
// 	var input string
// 	fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)
// 	fmt.Scan(&input)

// 	if question == input {
// 		fmt.Println("正解!")
// 		// ポインタを使って加算してください
// 		*scorePtr += 10

// 	} else {
// 		fmt.Println("不正解!")
// 	}
// }

/*
セクション7
*/
// package main

// 演習
/*
Q1. 以下のように文字列を連結させて出力したいコードがありますが、

test1!
test1!test2!
test1!test2!test3!
test1!test2!test3!test4!
以下のコードには間違いがあります。上記の出力になるように修正してください。

package main

import "fmt"

func goroutine(s []string, c chan int){
    sum := ""
    for _, v := range s{
        sum += v
        c <- sum
    }
}

func main(){
    words := []string{"test1!", "test2!", "test3!", "test4!"}
    c := make(chan int)
    go goroutine(words, c)
    for w := range c{
        fmt.Println(w)
    }
}
*/

// import "fmt"

// func goroutine(s []string, c chan string) {
// 	defer close(c)
// 	sum := ""
// 	for _, v := range s {
// 		sum += v
// 		c <- sum
// 	}
// }

// func main() {
// 	words := []string{"test1!", "test2!", "test3!", "test4!"}
// 	c := make(chan string)
// 	go goroutine(words, c)
// 	for w := range c {
// 		fmt.Println(w)
// 	}
// }

// 57.sync.Mutex
// type Counter struct {
// 	v   map[string]int
// 	mux sync.Mutex
// }

// func (c *Counter) Inc(key string) {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	c.v[key]++
// }

// func (c *Counter) Value(key string) int {
// 	c.mux.Lock()
// 	c.mux.Unlock()
// 	return c.v[key]
// }

// func main() {
// 	// c := make(map[string]int)
// 	c := Counter{v: make(map[string]int)}
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c.Inc("key")
// 		}
// 	}()
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c.Inc("key")
// 		}
// 	}()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(c, c.Value("key"))
// }

// 56.Default Selectionと for break
// func main() {
// 	tick := time.Tick(100 * time.Millisecond)
// 	boom := time.After(500 * time.Millisecond)
// 	// select for を抜ける
// OuterLoop:
// 	for {
// 		select {
// 		case <-tick:
// 			fmt.Println("tick.")
// 		case <-boom:
// 			fmt.Println("boom.")
// 			break OuterLoop
// 			// retrun
// 		default:
// 			fmt.Println(".")
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}
// 	fmt.Println("#######")
// }

// 55.channelとselect
// func goroutine1(ch chan string) {
// 	for {
// 		ch <- "packet from 1"
// 		time.Sleep(3 * time.Second)
// 	}
// }

// func goroutine2(ch chan string) {
// 	for {
// 		ch <- "packet from 2"
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)
// 	go goroutine1(c1)
// 	go goroutine2(c2)

// 	for {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println(msg1)
// 		case msg2 := <-c2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }

// 54.fan-out fan-in
// func producer(first chan int) {
// 	defer close(first)
// 	for i := 0; i < 10; i++ {
// 		first <- i
// 	}
// }

// // chanがoutかinか明示的にするパターン
// func multi2(first <-chan int, second chan<- int) {
// 	defer close(second)
// 	for i := range first {
// 		second <- i * 2
// 	}
// }

// func multi4(second chan int, third chan int) {
// 	defer close(third)
// 	for i := range second {
// 		third <- i * 4
// 	}
// }

// func main() {
// 	first := make(chan int)
// 	second := make(chan int)
// 	third := make(chan int)

// 	go producer(first)
// 	go multi2(first, second)
// 	go multi4(second, third)
// 	for result := range third {
// 		fmt.Println(result)
// 	}
// }

// 53.producerとconsumer
// func producer(ch chan int, i int) {
// 	// something
// 	ch <- i * 2
// }

// func consumer(ch chan int, wg *sync.WaitGroup) {
// 	for i := range ch {
// 		// エラー時にwg.Done()が呼ばれない対策としてインナーfunc()を利用
// 		func() {
// 			defer wg.Done()
// 			fmt.Println("process", i*1000)
// 		}()
// 	}
// 	fmt.Println("#######")
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	// Producer
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go producer(ch, i)
// 	}

// 	// Consumer
// 	go consumer(ch, &wg)
// 	wg.Wait()
// 	// close がないとfor文が完了しない
// 	close(ch)
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Done")
// }

// 52.channelのrangeとclose
// func goroutine(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 		fmt.Println(sum)
// 	}
// 	close(c)
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int, len(s))
// 	go goroutine(s, c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

// 51.Buffered Channels
// func main() {
// 	// 第二引数でバッファの数を指定
// 	ch := make(chan int, 2)
// 	ch <- 100
// 	fmt.Println(len(ch))
// 	ch <- 200
// 	fmt.Println(len(ch))
// 	// close で channelを閉じる。closeがないと、forループで3つ目も取りに行こうとしてエラーになる
// 	close(ch)

// 	for c := range ch {
// 		fmt.Println(c)
// 	}
// 	// 値を取り出した後はlen が0に戻る
// 	fmt.Println(len(ch))

// 	// 3つ目の受信はエラー
// 	// ch <- 300
// 	// fmt.Println(len(ch))
// }

// 50.channel
// func goroutine1(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// func goroutine2(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int) // 15, 15
// 	go goroutine1(s, c)
// 	go goroutine2(s, c)
// 	// 受信が終わるまで待機
// 	x := <-c
// 	fmt.Println(x)
// 	y := <-c
// 	fmt.Println(y)
// }

// 49.goroutine と sync.WaitGroup
// 並列処理
// func goroutine(s string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func nomal(s string) {
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	// 並列処理
// 	// 先にnomal()の処理が終了すると、gooroutine()が処理されないのでsync.WaitGroupを利用
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go goroutine("hello", &wg)
// 	nomal("world")
// 	wg.Wait()
// }

/**
セクション6
*/
// セクション6 Structオリエンテッド

// package main

/*
演習
Q1. 以下に、7と表示されるメソッドを作成してください。

package main

import (
    "fmt"
)

type Vertex struct{
    X, Y int
}

func main(){
    v := Vertex{3, 4}
    fmt.Println(v.Plus())
}
Q2 X is 3! Y is 4! と表示されるStringerを作成してください。

package main

import (
    "fmt"
)

type Vertex struct{
    X, Y int
}

func main(){
    v := Vertex{3, 4}
    fmt.Println(v)
}
*/

// Q1.
// type Vertex struct {
// 	X, Y int
// }

// func (v Vertex) Plus() int {
// 	return v.X + v.Y
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Plus())
// }

// Q2.
// type Vertex struct {
// 	X, Y int
// }

// func (v Vertex) String() string {
// 	return fmt.Sprintf("X is %d! Y is %d!", v.X, v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v)
// }

// 46.カスタムエラー
// type UserNotFound struct {
// 	UserName string
// }

// func (e *UserNotFound) Error() string {
// 	return fmt.Sprintf("User not found: %v", e.UserName)
// }

// func myFunc() error {
// 	ok := false
// 	if ok {
// 		return nil
// 	}
// 	return &UserNotFound{UserName: "Mike"}
// }

// func main() {
// 	if err := myFunc(); err != nil {
// 		fmt.Println(err)
// 	}
// }

// 45.Stringer
// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("My name is %v.", p.Name)
// }

// func main() {
// 	mike := Person{"Mike", 22}
// 	fmt.Println(mike)
// }

// 44.タイプアサーションとSwitch type文
// func do(i interface{}) {
// 	// intのみ
// 	// ii := i.(int)
// 	// ii *= 2
// 	// fmt.Println(ii)

// 	// stringのみ
// 	// ss := i.(string)
// 	// fmt.Println(ss + "!")

// 	// switch type文
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println(v * 2)
// 	case string:
// 		fmt.Println(v + "!")
// 	default:
// 		fmt.Printf("I don't know %T\n", v)
// 	}
// }

// func main() {
// 	do(10)
// 	do("Mike")
// 	do(true)
// }

// 43.インターフェイスとダックタイピング
// type Human interface {
// 	Say() string
// }

// type Person struct {
// 	Name string
// }

// func (p *Person) Say() string {
// 	p.Name = "Mr." + p.Name
// 	fmt.Println(p.Name)
// 	return p.Name
// }

// func DriveCar(human Human) {
// 	if human.Say() == "Mr.Mike" {
// 		fmt.Println("run")
// 	} else {
// 		fmt.Println("Get out")
// 	}
// }

// func main() {
// 	var mike Human = &Person{"Mike"}
// 	var x Human = &Person{"X"}

// 	// mike.Say()
// 	DriveCar(mike)
// 	DriveCar(x)

// }

// 41.Embedded（埋め込み）
// type Vertex struct {
// 	x, y int
// }

// type Vertex3D struct {
// 	Vertex
// 	z int
// }

// func (v Vertex3D) Area3D() int {
// 	return v.x * v.y * v.z
// }

// func (v *Vertex3D) Scale(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// 	v.z = v.z * i
// }

// func New(x, y, z int) *Vertex3D {
// 	return &Vertex3D{Vertex{x, y}, z}
// }

// func main() {
// 	v := New(3, 4, 5)
// 	v.Scale(10)
// 	fmt.Println(v.Area3D())

// }

// 40. コンストラクタ
/* pythonのclass例
class Vertex(object):
	def __init__(self, x, y):
		self._x = x
		self._y = y

	def area(self):
		return self._x * self._y

	def scale(self, i):
		self._x = self._x * i
		self._y = self._y * i

v = Vertex(3, 4)
v.scale(10)
print(v.area())
*/

// type Vertex struct {
// 	x, y int
// }

// // 値レシーバー（classの代わり）
// func (v Vertex) Area() int {
// 	return v.x * v.y
// }

// // ポインタレシーバー
// func (v *Vertex) Scale(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// }

// func New(x, y int) *Vertex {
// 	return &Vertex{x, y}
// }

// // func Area(v Vertex) int {
// // 	return v.x * v.y
// // }

// func main() {
// 	// v := Vertex{3, 4}
// 	// fmt.Println(Area(v))
// 	v := New(3, 4)
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// }

// 39. ポインタとポインタレシーバーと値レシーバー
// type Vertex struct {
// 	X, Y int
// }

// // 値レシーバー（classの代わり）
// func (v Vertex) Area() int {
// 	return v.X * v.Y
// }

// // ポインタレシーバー
// func (v *Vertex) Scale(i int) {
// 	v.X = v.X * i
// 	v.Y = v.Y * i
// }

// // func Area(v Vertex) int {
// // 	return v.X * v.Y
// // }

// func main() {
// 	v := Vertex{3, 4}
// 	// fmt.Println(Area(v))
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// }

/*
セクション5
*/
// package main

// セクション5 ポインタ
// 演習
// Q1. 以下の実行結果をコードを書かずに答えてください。エラーがおきますか？正しく表示されるとすると何が表示字されますか？
// A.10

// package main

// import "fmt"

// import (
// 	"fmt"
// )

// func main() {
// 	var i int = 10
// 	var p *int
// 	p = &i
// 	var j int
// 	j = *p
// 	fmt.Println(j)
// }

// Q2. 以下の実行結果をコードを書かずに答えてください。エラーがおきますか？正しく表示されるとすると何が表示されますか？

// package main

// func main() {
// 	var i int = 100
// 	var j int = 200
// 	var p1 *int
// 	var p2 *int
// 	p1 = &i
// 	p2 = &j
// 	i = *p1 + *p2

// 	p2 = p1
// 	j = *p2 + i
// 	fmt.Println(j)
// }

// 36.struct
// type Vertex struct {
// 	// 小文字だとスコープがprivateになる
// 	// X int
// 	// Y int
// 	X, Y int
// 	S    string
// }

// func changeVertex(v Vertex) {
// 	v.X = 1000
// }

// func changeVertex2(v *Vertex) {
// 	v.X = 1000
// 	// structでは同じ意味
// 	// (*v).X = 1000
// }

// func main() {
// 	v := Vertex{1, 2, "test"}
// 	changeVertex(v)
// 	fmt.Println(v)

// 	v2 := &Vertex{1, 2, "test"}
// 	changeVertex2(v2)
// 	fmt.Println(v2)
// 	fmt.Println(*v2)

// 	// v := Vertex{X: 1, Y: 2}
// 	// fmt.Println(v)
// 	// fmt.Println(v.X, v.Y)
// 	// v.X = 100
// 	// fmt.Println(v.X, v.Y)

// 	// v2 := Vertex{X: 1}
// 	// fmt.Println(v2)

// 	// v3 := Vertex{1, 2, "test"}
// 	// fmt.Println(v3)

// 	// v4 := Vertex{}
// 	// fmt.Println(v4)

// 	// var v5 Vertex
// 	// fmt.Println(v5)

// 	// // ポインタが返る
// 	// v6 := new(Vertex)
// 	// fmt.Println(v6)

// 	// v7 := &Vertex{}
// 	// fmt.Println(v7)
// }

// 35.newとmakeの違い
// func main() {

// 	s := make([]int, 0)
// 	fmt.Printf("%T\n", s)

// 	m := make(map[string]int)
// 	fmt.Printf("%T\n", m)

// 	ch := make(chan int)
// 	fmt.Printf("%T\n", ch)

// 	var p *int = new(int)
// 	fmt.Printf("%T\n", p)

// 	var st = new(struct{})
// 	fmt.Printf("%T\n", st)

// 	// メモリ確保
// 	// var p *int = new(int)
// 	// fmt.Println(p)
// 	// fmt.Println(*p)
// 	// *p++
// 	// fmt.Println(*p)

// 	// メモリ確保しない
// 	// var p2 *int
// 	// fmt.Println(p2)
// 	// // メモリ確保しないのでエラー
// 	// *p2++
// 	// fmt.Println(p2)
// }

//34.ポインタ
// func one(x int) {
// 	x = 1
// }

// func one2(x *int) {
// 	*x = 1
// }

// func main() {
// 	var n int = 100
// 	one(n)
// 	// 値は変わらない
// 	fmt.Println(n)

// 	one2(&n)
// 	fmt.Println(n)

// 	// fmt.Println(&n)

// 	// var p *int = &n
// 	// fmt.Println(p)

// 	// fmt.Println(*p)
// }

/*
セクション4
*/
// package main

// セクション4 ステートメント
// 32.演習
/*
Q1 . 以下のスライスから一番小さい数を探して出力するコードを書いてください。

l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

Q2. 以下の果物の価格の合計を出力するコードを書いてください。

m := map[string]int{
    "apple":  200,
    "banana": 300,
    "grapes": 150,
    "orange": 80,
    "papaya": 500,
    "kiwi":   90,
}
*/

// Q1
// func main() {
// 	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

// 	var min int
// 	for i, v := range l {
// 		if i == 0 {
// 			min = v
// 			continue
// 		}

// 		if min > v {
// 			min = v
// 		}
// 	}
// 	fmt.Println(min)
// }

// Q.2
// func main() {
// 	m := map[string]int{
// 		"apple":  200,
// 		"banana": 300,
// 		"grapes": 150,
// 		"orange": 80,
// 		"papaya": 500,
// 		"kiwi":   90,
// 	}

// 	sum := 0
// 	for _, v := range m {
// 		sum += v
// 	}
// 	fmt.Println(sum)
// }

// 31.パニックとリカバー
// panic を自分がコーディングするのは推奨されていない。何が起こったかわからないという状況を作らないというGoの思想
// func thirdPartyConnectDB() {
// 	panic("Unable to connect database")
// }

// func save() {
// 	defer func() {
// 		s := recover()
// 		fmt.Println(s)
// 	}()
// 	thirdPartyConnectDB()
// }

// func main() {
// 	save()
// 	fmt.Println("OK?")
// }

// 30.エラーハンドリング（例外の代わり）
// func main() {
// 	// 返り値が2つ
// 	file, err := os.Open("./section4.go")
// 	if err != nil {
// 		log.Fatalln("Error!")
// 	}
// 	defer file.Close()
// 	data := make([]byte, 100)
// 	count, err := file.Read(data)
// 	if err != nil {
// 		log.Fatalln("Error2!")
// 	}
// 	fmt.Println(count, string(data))

// 	// 返り値が1つなら1行で書く
// 	if err = os.Chdir("test"); err != nil {
// 		log.Fatalln("Error")
// 	}
// }

// // 29.log
// func LoggingSettings(logFile string) {
// 	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	multiLogFile := io.MultiWriter(os.Stdout, logfile)
// 	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
// 	log.SetOutput(multiLogFile)
// }

// func main() {
// 	LoggingSettings("test.log")
// 	_, err := os.Open("faeifjojdf")
// 	if err != nil {
// 		log.Fatalln("Exit!", err)
// 	}

// 	log.Println("logging")
// 	log.Printf("%T  %v", "test", "test")

// 	// fatalはexitする
// 	log.Fatalf("%T, %v", "test", "test")
// 	log.Fatalln("error!")
// 	fmt.Println("ok!")
// }

// 28.defer
// func foo() {
// 	defer fmt.Println("world foo")

// 	fmt.Println("hello foo")
// }

// func main() {
// 	// defer fmt.Println("world")

// 	// foo()

// 	// fmt.Println("hello")

// 	// fmt.Println("run")
// 	// defer fmt.Println("1")
// 	// defer fmt.Println("2")
// 	// defer fmt.Println("3")
// 	// fmt.Println("success")

// 	// deferの実用例
// 	file, _ := os.Open("./section4.go")
// 	defer file.Close()
// 	data := make([]byte, 100)
// 	file.Read(data)
// 	fmt.Println(string(data))
// }

// 27.switch
// func getOsName() string {
// 	return "deeefault"
// }

// func main() {
// 	// os := "win"
// 	// 1文でも書ける ※スコープはswitch文の中だけ
// 	switch os := getOsName(); os {
// 	case "mac":
// 		fmt.Println("mac!")
// 	case "windows":
// 		fmt.Println("windows!")
// 	// defaultはなくてもOK
// 	default:
// 		fmt.Println("default!", os)
// 	}

// 	// switchの条件を書かない
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("afternoon!")
// 	}
// }

// 26.range
// func main() {
// 	l := []string{"python", "go", "ruby"}
// 	for i := 0; i < len(l); i++ {
// 		fmt.Println(i, l[i])
// 	}

// 	// 違う書き方
// for i, v := range l {
// 	fmt.Println(i, v)
// }

// 	// iを使用しない
// 	for _, v := range l {
// 		fmt.Println(v)
// 	}

// 	// map
// 	m := map[string]int{"apple": 100, "banana": 200}
// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}

// 	for k := range m {
// 		fmt.Println(k)
// 	}

// 	for _, v := range m {
// 		fmt.Println(v)
// 	}
// }

// 25.for
// func main() {
// 	for i := 0; i < 10; i++ {
// 		if i == 3 {
// 			fmt.Println("continue")
// 			continue
// 		}

// 		if i > 5 {
// 			fmt.Println("break")
// 			break
// 		}
// 		fmt.Println(i)
// 	}

// 	sum := 1
// 	for sum < 10 {
// 		sum += sum
// 		fmt.Println(sum)
// 	}
// 	fmt.Println(sum)
// }

// 24.if
// func by2(num int) string {
// 	if num%2 == 0 {
// 		return "ok"
// 	} else {
// 		return "no"
// 	}
// }

// func main() {
// 	result := by2(10)
// 	if result == "ok" {
// 		fmt.Println("great")
// 	}

// 	if result2 := by2(10); result2 == "ok" {
// 		fmt.Println("great 2")
// 	}

// 	// num := 11
// 	// if num%2 == 0 {
// 	// 	fmt.Println("by 2")
// 	// } else if num%3 == 0 {
// 	// 	fmt.Println("by 3")
// 	// } else {
// 	// 	fmt.Println("else")
// 	// }

// 	x, y := 11, 12
// 	if x == 10 && y == 10 {
// 		fmt.Println("&&")
// 	}

// 	if x == 10 || y == 10 {
// 		fmt.Println("||")
// 	}
// }

/*
セクション3
*/
// package main

// セクション3 定義
// セクション3 演習
/*
演習
Q1. 以下の1.11をint型に変換して出力してください。

f := 1.11

Q2. コードを書かずに以下の出力結果を答えてください。

s := []int{1, 2, 5, 6, 2, 3, 1}
fmt.Println(s[2:4])
Q3. 以下のコードを実行した時に

fmt.Printf("%T %v", m, m)

以下のような出力結果となるmを作成してください。

map[string]int map[Mike:20 Nancy:24 Messi:30]
*/

// func main() {
// 	// Q1
// 	f := 1.11
// 	i := int(f)
// 	fmt.Printf("%T %v\n", i, i)

// 	// Q2
// 	s := []int{1, 2, 5, 6, 2, 3, 1}
// 	fmt.Println(s[2:4])

// 	// Q3
// 	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
// 	fmt.Printf("%T %v", m, m)
// }

// 21.可変長変数
// func foo(params ...int) {
// 	fmt.Println(len(params), params)
// 	for _, param := range params {
// 		fmt.Println(param)
// 	}
// }
// func main() {
// 	foo()
// 	foo(10)
// 	foo(10, 20)
// 	foo(10, 20, 30)

// 	// スライス
// 	s := []int{1, 2, 3}
// 	fmt.Println(s)

// 	foo(s...)
// }

// 20.クロージャ
// func incrementGenerator() func() int {
// 	x := 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// func circleArea(pi float64) func(radis float64) float64 {
// 	return func(radis float64) float64 {
// 		return pi * radis * radis
// 	}
// }

// func main() {
// 	counter := incrementGenerator()
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// 	fmt.Println(counter())

// 	c1 := circleArea(3.14)
// 	fmt.Println(c1(2))
// 	fmt.Println(c1(3))

// 	c2 := circleArea(3)
// 	fmt.Println(c2(2))
// 	fmt.Println(c2(3))
// }

// 19.関数
// func add(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	r := add(10, 20)
// 	fmt.Println(r)
// }

// func add(x, y int) (int, int) {
// 	return x + y, x - y
// }

// func cal(price, item int) (result int) {
// 	result = price * item
// 	return // result
// }

// func main() {
// 	r1, r2 := add(10, 20)
// 	fmt.Println(r1, r2)

// 	r3 := cal(100, 2)
// 	fmt.Println(r3)

// 	f := func(x int) {
// 		fmt.Println("inner func", x)
// 	}
// 	f(1)

// 	func(x int) {
// 		fmt.Println("inner func", x)
// 	}(1)
// }

// 18. バイト型
// func main() {
// 	b := []byte{72, 73}
// 	fmt.Println(b)
// 	fmt.Println(string(b))

// 	c := []byte("HI")
// 	fmt.Println(c)
// 	fmt.Println(string(c))
// }

// 17.map
// func main() {
// 	m := map[string]int{"apple": 100, "banana": 200}
// 	fmt.Println(m)
// 	fmt.Println(m["apple"])
// 	m["banana"] = 300
// 	fmt.Println(m)
// 	m["new"] = 500
// 	fmt.Println(m)

// 	fmt.Println(m["nothing"])

// 	v, ok := m["apple"]
// 	fmt.Println(v, ok)

// 	v2, ok2 := m["nothing"]
// 	fmt.Println(v2, ok2)

// 	m2 := make(map[string]int)
// 	m2["pc"] = 5000
// 	fmt.Println(m2)

// 	// nil メモリ上にないので
// 	// var m3 map[string]int
// 	// m3["pc"] = 5000
// 	// fmt.Println(m3)

// 	var s []int
// 	if s == nil {
// 		fmt.Println("nil")
// 	}
// }

// 16.スライスのmakeとcap
// func main() {
// n := make([]int, 3, 5)
// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
// n = append(n, 1, 3)
// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
// n = append(n, 4, 5, 6)
// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

// // lenとcapどちらも指定
// a := make([]int, 3)
// fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

// b := make([]int, 0) // 0のスライスをメモリに確保
// var c []int // nill メモリ確保しない
// fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
// fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

// c = make([]int, 5)
// for i := 0; i < 5; i++ {
// 	c = append(c, i)
// 	fmt.Println(c)
// }
// fmt.Println(c)

// c = make([]int, 0, 5)
// for i := 0; i < 5; i++ {
// 	c = append(c, i)
// 	fmt.Println(c)
// }
// fmt.Println(c)

// }

// 15. スライス
// func main() {
// 	n := []int{1, 2, 3, 4, 5}
// 	fmt.Println(n)
// 	fmt.Println(n[3])
// 	fmt.Println(n[2:4])
// 	fmt.Println(n[:2])
// 	fmt.Println(n[2:])
// 	fmt.Println(n[:])

// 	n[2] = 100
// 	fmt.Println(n)

// 	var board = [][]int{
// 		[]int{0, 1, 2},
// 		[]int{3, 4, 5},
// 		[]int{6, 7, 8},
// 	}
// 	fmt.Println(board)

// 	n = append(n, 100, 200, 300)
// 	fmt.Println(n)
// }

// 14.配列
// func main() {
// 	var a [2]int
// 	a[0] = 100
// 	a[1] = 200
// 	fmt.Println(a)

// 	// 配列はサイズを変更できない
// 	// var b [2]int = [2]int{100, 200}
// 	// // b = append(b, 300) // 配列は値の追加ができない
// 	// fmt.Println(b)

// 	// slice
// 	var b []int = []int{200, 300}
// 	b = append(b, 300)
// 	fmt.Println(b)
// }

// 13.型変換
// func main() {
// 	var x int = 1
// 	xx := float64(x)
// 	fmt.Printf("%T %v %f\n", xx, xx, xx)

// 	var y float64 = 2.1
// 	yy := int(y)
// 	fmt.Printf("%T %v %d\n", yy, yy, yy)

// 	var s string = "14"
// 	// stringをint型に単純には型変換はできない
// 	// i := int(s)
// 	i, _ := strconv.Atoi(s)
// 	fmt.Printf("%T, %v\n", i, i)

// 	h := "Hello world"
// 	fmt.Println(h[0], string(h[0]))
// }

// 12.論理値型
// func main() {
// 	t, f := true, false
// 	fmt.Printf("%T %v %t\n", t, t, t)
// 	fmt.Printf("%T %v %t\n", f, f, 0)

// 	fmt.Println(true && true)
// 	fmt.Println(true && false)
// 	fmt.Println(false && false)

// 	fmt.Println(true || true)
// 	fmt.Println(true || false)
// 	fmt.Println(false || false)

// 	fmt.Println(!true)
// 	fmt.Println(!false)

// }

// 11.文字列型
// func main() {
// 	fmt.Println("Hello world")
// 	fmt.Println("Hello" + "world")
// 	fmt.Println("Hello world"[0])
// 	fmt.Println(string("Hello world"[0]))
// 	fmt.Printf("%T", "Hello"[0])

// 	var s string = "Hello world"
// 	fmt.Println(strings.Replace(s, "H", "X", 1))
// 	fmt.Println(strings.Contains(s, "world"))

// 	fmt.Println("test\n" +
// 		"test")
// 	fmt.Println(`test
// 	test
// 	      test`)

// 	fmt.Println("\"")
// 	fmt.Println(`"`)
// }

// 10.数値型
// func main() {
// var (
// 	u8  uint8     = 255
// 	i8  int8      = 127
// 	f32 float32   = 0.2
// 	c64 complex64 = -5 + 12i
// )

// fmt.Println(u8, i8, f32, c64)

// fmt.Printf("type=%T, value=%v", u8, u8)

// fmt.Println("1 + 1 =", 1+1)
// fmt.Println("10 - 1 =", 10-1)
// fmt.Println("10 / 2 =", 10/2)
// fmt.Println("10 / 3 =", 10/3)
// fmt.Println("10.0 / 3 =", 10.0/3)
// fmt.Println("10 / 3.0 =", 10/3.0)
// fmt.Println("10 % 2 =", 10%2)
// fmt.Println("10 % 3 =", 10%3)

// x := 0
// x++
// fmt.Println(x)
// x--
// fmt.Println(x)

// 	fmt.Println(1 << 0) // 0001
// 	fmt.Println(1 << 1) // 0010
// 	fmt.Println(1 << 2) // 0100
// 	fmt.Println(1 << 3) // 1000

// }

// 定数
// const Pi = 3.14

// const (
// 	Username = "test_user"
// 	Password = "password"
// )

// var big int = 9223372036854775807 + 1 // オーバーフロー
// const Big = 9223372036854775807 + 1

// func main() {
// 	fmt.Println(Pi, Username, Password)
// 	// fmt.Printf("%T", Pi)
// 	fmt.Println(Big - 1)
// }

// func init() {
// 	fmt.Println("init")
// }

// func bazz() {
// fmt.Println("bazz")
// }

// var (
// 	i    int     = 1
// 	f64  float64 = 1.2
// 	s    string  = "test"
// 	t, f bool    = true, false
// )

// func foo() {
// 	xi := 1
// 	xi = 2
// 	xf64 := 1.2
// 	xs := "test"
// 	xt, xf := true, false
// 	fmt.Println(xi, xf64, xs, xt, xf)
// 	fmt.Printf("%T\n", xf64)
// 	fmt.Printf("%T\n", xi)
// }

// func main() {
// 	fmt.Println(i, f64, s, t, f)
// 	foo()

// 	// bazz()
// 	// fmt.Println("hello world", time.Now())
// 	// fmt.Println(user.Current())
// }

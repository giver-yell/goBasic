package main

// 75.hmac でAPI認証

// 74.json.Unmarshal と marshal と encode
// type T struct{}

// type Person struct {
// 	Name      string   `json:"name"`
// 	Age       int      `json:"age,omitempty"`
// 	NickNames []string `json:"nicknames"`
// 	T         *T       `json:"T,omitempty"`
// }

// // Unmarshalのカスタマイズ
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

// // Marshalのカスタマイズ
// // func (p Person) MarshalJSON() ([]byte, error) {
// // 	v, err := json.Marshal(&struct {
// // 		Name string
// // 	}{
// // 		Name: "Mr." + p.Name,
// // 	})
// // 	return v, err
// // }

// func main() {
// 	b := []byte(`{"name":"mike", "age":0,"nicknames":["a","b","c"]}`)
// 	var p Person
// 	if err := json.Unmarshal(b, &p); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(p.Name, p.Age, p.NickNames)

// 	v, _ := json.Marshal(p)
// 	fmt.Println(string(v))
// }

// 73.http
// func main() {
// 	// resp, _ := http.Get("http://example.com")
// 	// defer resp.Body.Close()
// 	// body, _ := ioutil.ReadAll(resp.Body)
// 	// fmt.Println(string(body))

// 	base, err := url.Parse("http://example.com/abdeojo")
// 	reference, _ := url.Parse("/test?a=1&b=2")
// 	endpoint := base.ResolveReference(reference).String()
// 	fmt.Println(base, err)
// 	fmt.Println(endpoint)
// 	req, _ := http.NewRequest("GET", endpoint, nil)
// 	// req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
// 	req.Header.Add("If-None-Match", `W/"wyzzy"`)
// 	q := req.URL.Query()
// 	q.Add("c", "3&%")
// 	fmt.Println(q)
// 	fmt.Println(q.Encode())
// 	req.URL.RawQuery = q.Encode()

// 	var client *http.Client = &http.Client{}
// 	resp, _ := client.Do(req)
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(body))
// }

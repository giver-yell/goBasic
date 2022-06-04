package main

import (
	"fmt"
	"sort"
)

// 69.Sort
func main() {
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)
}

// 68.regex
// func main() {
// 	match, _ := regexp.MatchString("a([a-z0-9]+)e", "appl0e")
// 	fmt.Println(match)

// 	r := regexp.MustCompile("a([a-z]+)e")
// 	ms := r.MatchString("apple")
// 	fmt.Println(ms)

// 	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-z0-9]+)$")
// 	fs := r2.FindString("/view/test")
// 	fmt.Println(fs)
// 	fss := r2.FindStringSubmatch("/view/test")
// 	fmt.Println(fss, fss[0], fss[1], fss[2])
// }

// 67.time
/*
https://pkg.go.dev/time

Constants ¶
View Source
const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)
*/

// func main() {
// 	t := time.Now()
// 	fmt.Println(t)
// 	// postgreではRFC3339が標準
// 	fmt.Println(t.Format(time.RFC3339))
// 	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
// }

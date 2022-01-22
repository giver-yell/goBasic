// 62.testing
package mylib

import "testing"

// var Debug bool = true

func TestAvarage(t *testing.T) {
	// if Debug {
	// 	t.Skip("Skip reason")
	// }
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}

package avlogparser

import (
	"fmt"
)

// TestDecoding : ril
func TestDecoding() {
	test := "dGVzdA=="
	fmt.Println("testing: " + test)
	fmt.Println("Decoded: " + DecodeToString(test))
}

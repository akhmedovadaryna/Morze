package logic

import (
	"testing"
	"fmt"
)

func TestFormat_text(t *testing.T) {
	var m string
	m = format_text("AAAAA.....AAAAA")
	fmt.Printf(m)
	if m != "AAAAAAAAAA"{
		t.Error("Expected AAAAAAAAAA, got", m)
	}
}
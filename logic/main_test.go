package Morze

import (
	"testing"
	"fmt"
)

func TestMorze(t *testing.T) {
	var m string
	m = Morze("--._---____--._._-_")
	fmt.Printf(m)
	if m != "GO GET"{
		t.Error("Expected GO GET, got", m)
	}
}

func TestTextToMorze(t *testing.T) {
	var ttm string
	ttm = TextToMorze("Go")
	if ttm != "--._---_\n" {
		t.Error("Expecteg --._---_, got", ttm)
	}
}

func BenchmarkMorze(b *testing.B) {
	for n := 0; n < b.N; n++{
		Morze("--._---____--._._-_")
	}
}
func BenchmarkTextToMorze(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TextToMorze("GO GET")

	}
}